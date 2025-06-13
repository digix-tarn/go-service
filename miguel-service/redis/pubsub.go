package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	// "log"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func Subscribe(channel string, handleMessage func(msg string)) {
	pubsub := Rdb.Subscribe(Ctx, channel)
	ch := pubsub.Channel()
	go func() {
		for msg := range ch {
			handleMessage(msg.Payload)
		}
	}()
}

func Publish(channel string, msg string) error {
	return Rdb.Publish(Ctx, channel, msg).Err()
}
