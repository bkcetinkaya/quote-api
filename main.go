package main

import (
	"daily-quote/db"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	quote := db.ConnectToDatabase()

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
