package main

import (
	"flag"
	"log"
	"net/http"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var redisAddress = flag.String("redis-address", ":6379", "Address to the Redis server")

func main() {
	_, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = redigo.Dial("tcp", *redisAddress)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = http.Get("http://0.0.0.0:9200")
	if err != nil {
		log.Fatalln(err)
	}

}

// simple addition math func
func add(a, b int) int {
	return a + b
}
