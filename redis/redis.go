package redis

import (
	"log"
	"time"

	goredis "github.com/xuyu/goredis"
)

//New initialize redis connection from program
//127.0.0.1:6379 is default connection for local redis
func New() *goredis.Redis {
	address := "127.0.0.1:6379"

	client, err := goredis.Dial(&goredis.DialConfig{
		Address: address,
		Timeout: 10 * time.Second,
		MaxIdle: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[REDIS] Redis Connected with", address)
	return client
}
