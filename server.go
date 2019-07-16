// server.go

package main

import (
	"net/http"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()

	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// // CORS
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))

	// // Server
	// e.Run(standard.New(":1323"))

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
