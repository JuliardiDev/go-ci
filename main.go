package main

import (
	"log"
	"net/http"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var (
		redisConnection         = "Redis OK"
		dbConnection            = "DB OK"
		elasticSearchConnection = "Elasticsearch OK"
	)
	_, err := redigo.Dial("tcp", "redis:6379")
	if err != nil {
		redisConnection = "Failed: redis " + err.Error()
		log.Fatalln(redisConnection)
	}

	_, err = sqlx.Connect("postgres", "user=postgres dbname=postgres password=root host=db sslmode=disable")
	if err != nil {
		dbConnection = "Failed: db " + err.Error()
		log.Fatalln(dbConnection)
	}

	_, err = http.Get("http://es:9200")
	if err != nil {
		elasticSearchConnection = "Failed: elasticsearch " + err.Error()
		log.Fatalln(elasticSearchConnection)
	}

	log.Println(redisConnection)
	log.Println(dbConnection)
	log.Println(elasticSearchConnection)

}

// simple addition math func
func add(a, b int) int {
	return a + b
}
