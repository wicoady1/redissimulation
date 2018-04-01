package handler

import (
	"github.com/wicoady1/redissimulation/redis"
	"github.com/xuyu/goredis"
)

type Module struct {
	Redis *goredis.Redis
}

func InitModule() *Module {
	redis := redis.New()

	module := &Module{
		Redis: redis,
	}

	return module
}
