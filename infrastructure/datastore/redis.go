package datastore

import (
	"context"
	"time"

	r "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisClient(host map[string]string, pass string, db int) (client *r.Ring) {

	client = r.NewRing(&r.RingOptions{
		Addrs:        host,
		Password:     pass,
		DB:           db,
		DialTimeout:  time.Duration(30) * time.Second,
		WriteTimeout: time.Duration(30) * time.Second,
		ReadTimeout:  time.Duration(30) * time.Second,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return
}
