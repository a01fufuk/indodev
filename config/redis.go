package config

import (
	"flag"
	"fmt"

	"github.com/go-redis/redis"
)

var connection redis.Client

func ConnectRedisCloud() *redis.Client {
	connStr := fmt.Sprintf("%s:%s", REDISHost, REDISPort)
	var addr = flag.String("Server Cloud", connStr, "Redis server address")
	fmt.Println("Successful Connected to Redis Cloud:", string(*addr))

	rdb := redis.NewClient(&redis.Options{
		Addr:     *addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
