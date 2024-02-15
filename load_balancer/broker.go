package loadbalancer

import (
	"github.com/redis/go-redis/v9"
)

type Broker struct {
	address string
	redis   *redis.Client
}

func NewBroker(address string) *Broker {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Broker{
		address: address,
		redis:   rdb,
	}
}

type BrokerPayload struct {
	ServerAddress string `json:"serverAddress"`
}
