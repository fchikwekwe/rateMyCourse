// server.go

package main

import (
	// Standard library imports
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	// Gorm and postgres imports
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

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

// User Model
type User struct {
	gorm.Model
	email     string `gorm:"type:varchar(100);unique;not_null;primary_key"`
	username  string `gorm:"type:varchar(50);unique;not null"`
	firstName string `gorm:"type:varchar(30)`
	lastName  string `gorm:"type:varchar(30)`
	password  string `gorm:"type:varchar(30)`
}

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
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

	// Open the user database
	userDB := initUserDB()

	// GET Routes; showing templates
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.GET("/signup", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
