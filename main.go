package main

import (
	"fmt"
	"log"
	"net/http"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type logging struct {
}

func (lg logging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		redisConnection         = "Redis OK"
		dbConnection            = "DB OK"
		elasticSearchConnection = "Elasticsearch OK"
	)
	log.Println("I hitted")
	_, err := redigo.Dial("tcp", "redis:6379")
	if err != nil {
		redisConnection = "Failed: redis " + err.Error()
	}

	_, err = sqlx.Connect("postgres", "user=postgres dbname=postgres password=root host=db sslmode=disable")
	if err != nil {
		dbConnection = "Failed: db " + err.Error()
	}

	_, err = http.Get("http://elasticsearch:9200")
	if err != nil {
		elasticSearchConnection = "Failed: elasticsearch " + err.Error()
	}

	fmt.Fprintln(w, redisConnection)
	fmt.Fprintln(w, dbConnection)
	fmt.Fprintln(w, elasticSearchConnection)
}

func main() {
	lg := logging{}

	err := http.ListenAndServe(":8080", lg)
	fmt.Println(err)

}

// simple addition math func
func add(a, b int) int {
	return a + b
}
