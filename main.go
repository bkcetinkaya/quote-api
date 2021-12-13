package main

import (
	"daily-quote/db"
	util "daily-quote/utils"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	documentAmount := db.ConnectToDatabase()
	util.GetRandomInt(documentAmount)

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "blabla")
	})

	e.Logger.Fatal(e.Start(":8080"))

}
