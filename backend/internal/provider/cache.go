package provider

import (
	"cchart/internal/kernel"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func NewCache(config *kernel.Config) *redis.Client {
	c := config.Cache

	return redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + strconv.Itoa(c.Port),
		Password: c.Password,
		DB:       c.DB,
	})
}
