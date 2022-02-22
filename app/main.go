package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func Index(cache Cache) http.HandlerFunc {

	type Profile struct {
		Name  string `json:"name"`
		Image string `json:"image"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)

		query := r.URL.Query()
		username := query.Get("username")

		p := &Profile{}
		exists, err := cache.Get(username, p)

		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if exists {
			log.Println("data from cache")
			json.NewEncoder(rw).Encode(p)
			return
		}

		gp, err := requestGithubProfile(username)

		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		p = &Profile{
			Name:  gp.Name,
			Image: gp.Avatar,
		}

		log.Println("data from github api")
		cache.Set(username, p)
		json.NewEncoder(rw).Encode(p)
	}

}

func run() error {
	port := os.Getenv("PORT")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")

	redis := NewRedisCache(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
	}, 10)

	http.HandleFunc("/", Index(redis))

	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
