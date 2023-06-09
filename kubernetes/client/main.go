package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func redisCounter(w http.ResponseWriter, r *http.Request) {
	if "/favicon.ico" == r.URL.Path {
		return
	}
	log.Print("request: " + r.URL.Path + " ")

	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_URL")})
	client.Incr("kcount")

	val, err := client.Get("kcount").Result()
	if err != nil {
		panic(err)
	}
	result := string("key count: " + string(val))

	log.Println(result)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
	return
}

func main() {
	log.Print("Serving at 0.0.0.0:9090... ")
	http.HandleFunc("/", redisCounter)
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", nil))
}
