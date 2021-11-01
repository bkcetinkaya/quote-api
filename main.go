package main

import (
	"daily-quote/db/models"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "123456"
	dbname   = "quote-app"
)

func main() {
	app := fiber.New()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {

	})

	app.Listen(":8080")
}

func GetHandler(db *sql.DB) []models.Quote {
	rows, err := db.Query("SELECT * FROM quotes")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	quotes := []models.Quote{}

	for rows.Next() {
		var q models.Quote
		if err := rows.Scan(&q.ID, &q.Author, &q.Quote); err != nil {
			panic(err)
		}
		quotes = append(quotes, q)
	}

	return quotes
}
