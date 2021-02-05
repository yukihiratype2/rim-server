package event

import (
	"github.com/go-redis/redis/v8"
)

var event *redis.Client

func Start() {
	event = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
