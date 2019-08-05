// server.go

package main

import (
	"html/template"
	"io"
	"net/http"

	// // Import gorm and postgres
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"

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
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Basic Auth
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// maybe check if username is in database and then check is password is correct
		if username == "faith" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	// Template Renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "./public/index.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})

	// e.GET("/", root).Name = "example"

	// func root(c echo.Context) error {
	// 	return c.Render(http.StatusOk, "index.html", map[string]interface{}{
	// 		"name": "Dolly",
	// 	})
	// }

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
