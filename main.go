package main

import (
	"daily-quote/db"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	port := os.Getenv("PORT")

	e.GET("/", func(c echo.Context) error {

		quote := db.ConnectToDatabase()
		return c.JSON(http.StatusOK, quote)
	})

	e.Logger.Fatal(e.Start(":" + port))

}
