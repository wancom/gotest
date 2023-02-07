package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const PubSubCh = "test"

type Redis struct {
	client *redis.Client
}

func NewRedis() Redis {
	return Redis{client: redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})}
}

func (r *Redis) Publish(text string) error {
	ctx := context.Background()
	return r.client.Publish(ctx, PubSubCh, text).Err()
}

func (r *Redis) SubscribeCh(ctx context.Context) <-chan string {
	ch := make(chan string, 10)
	pubsub := r.client.Subscribe(ctx, PubSubCh)
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				break
			}
			ch <- msg.Payload
		}
	}()
	return ch
}
