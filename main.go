package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	_, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=roo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
}

// simple addition math func
func add(a, b int) int {
	return a + b
}
