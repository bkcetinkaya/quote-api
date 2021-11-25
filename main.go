package main

import (
	"context"
	db "daily-quote/db/sqlc"
	util "daily-quote/utils"
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries

func main() {

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, GetQuote())
	})

	e.Logger.Fatal(e.Start(":8080"))

}

func GetQuote() db.Quote {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err2 := sql.Open(config.DBDriver, config.DBSource)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer conn.Close()

	testQueries = db.New(conn)

	q, err3 := testQueries.GetQuote(context.Background(), int64(GetRandomInt()))
	if err3 != nil {
		panic(err3)
	}

	return q
}

func GetRandomInt() int {

	rowCount, err := testQueries.GetQuoteRows(context.Background())
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	num := 1 + rand.Intn(int(rowCount+1)-1)
	return num
}
