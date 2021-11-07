package main

/* -

- Random int dönderen fonksiyonun max değeri databasedeki toplam quote sayısından almalı. (query.sql dosyasına getQuotes ekle)




*/

import (
	"context"
	db "daily-quote/db/sqlc"
	"database/sql"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:123456@localhost:5432/quote-app?sslmode=disable"
)

var testQueries *db.Queries

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(GetQuote())
	})
	app.Listen(":8080")
}

func GetQuote() db.Quote {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}

	testQueries = db.New(conn)

	q, err2 := testQueries.GetQuote(context.Background(), int64(GetRandomInt()))
	if err2 != nil {
		panic(err)
	}

	//data, _ := json.Marshal(q)
	//quote := string(data)
	return q
}

func GetRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	num := 1 + rand.Intn(4-1)
	return num
}
