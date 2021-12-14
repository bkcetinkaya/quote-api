package main

import (
	"daily-quote/db"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {

		quote := db.ConnectToDatabase()
		return c.JSON(http.StatusOK, quote)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
