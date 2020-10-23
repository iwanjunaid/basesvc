package datastore

import (
	"time"

	r "github.com/go-redis/redis"
)

const (
	CfgRedisHost = "database.redis.master.address"
	CfgRedisPass = "database.redis.master.password"
	CfgRedisDB   = "database.redis.master.db"
)

func NewRedisClient(host []string, pass string, db int) (client *r.Client) {

	if len(host) > 1 {
		// @TODO: using cluster
	}

	client = r.NewClient(&r.Options{
		Addr:         host[0],
		Password:     pass,
		DB:           db,
		DialTimeout:  time.Duration(30) * time.Second,
		WriteTimeout: time.Duration(30) * time.Second,
		ReadTimeout:  time.Duration(30) * time.Second,
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	return
}
