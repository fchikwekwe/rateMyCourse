// server.go

package main

import (
	// Standard library imports
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	// Gorm and postgres imports
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"

	// Echo and echo middleware imports
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

// User stores user data
type User struct {
	gorm.Model
	createdAt time.Time
	updatedAt time.Time
	userID    int    `gorm:"type:int;primary_key"`
	email     string `gorm:"type:varchar(100);unique;not_null"`
	username  string `gorm:"type:varchar(50);unique;not null"`
	firstName string `gorm:"type:varchar(30);not_null"`
	lastName  string `gorm:"type:varchar(30);not_null"`
	password  []byte `gorm:"type:varchar(30);not_null"`
	reviews   []Review
}

func initUserDB() *gorm.DB {
	// Open DB
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("DB_PORT"), os.Getenv("user"), os.Getenv("dbname")))
	// Throw error if connection fails
	if err != nil {
		log.Print(err)
	}
	// Automigrate the DB
	db.AutoMigrate(&User{})

	return db
}

type Review struct {
	gorm.Model
	createdAt *time.Time
	updatedAt *time.Time
	reviewID  int    `gorm:"type:int;primary_key"`
	rating    string `gorm:"type:varchar(30);not_null"`
	text      string `gorm:"type:text"`
}

func initReviewDB() *gorm.DB {
	// Open DB
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("DB_PORT"), os.Getenv("user"), os.Getenv("dbname")))
	// Throw error if connection fails
	if err != nil {
		log.Print(err)
	}
	// Automigrate the DB
	db.AutoMigrate(&Review{})

	return db
}

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

func signup(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Decode the request body and create a new `Credentials` instance
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)

	if err != nil {
		// return an error message if there is something wrong with the request body
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Salt and hash the password with the bcrypt algorithm
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	// Next, insert the username along with hashed password into the db
	if _, err = db.Query(&User{ // what does db.Create return?
		createdAt: time.Now(),
		updatedAt: time.Now(),
		email:     email,
		username:  creds.Username,
		firstName: firstName,
		lastName:  lastName,
		password:  hashedPassword,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

/*
	email     string `gorm:"type:varchar(100);unique;not_null"`
	username  string `gorm:"type:varchar(50);unique;not null"`
	firstName string `gorm:"type:varchar(30);not_null"`
	lastName  string `gorm:"type:varchar(30);not_null"`
	password  string `gorm:"type:varchar(30);not_null"`
	reviews   []Review
*/

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("./build"))

	// Template Renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./build/*.html")),
	}
	e.Renderer = renderer

	// Open the user and review databases
	userDB := initUserDB()
	reviewDB := initReviewDB()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.GET("/signup", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.POST("/signup", signup(stuff, otherStuff, userDB))

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
