package redis

import (
	"context"
	"time"

	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
)

type InternalRedisImpl struct {
	rdb   *redis.Ring
	cache *cache.Codec
}

type InternalRedis interface {
	Get(ctx context.Context, key string) (resp interface{}, err error)
	Create(ctx context.Context, key string, value interface{}) error
}

func (ar *InternalRedisImpl) Get(ctx context.Context, key string) (resp interface{}, err error) {
	if err := ar.cache.Get(key, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (ar *InternalRedisImpl) Create(ctx context.Context, key string, value interface{}) error {
	if err := ar.cache.Set(&cache.Item{
		Ctx:    ctx,
		Key:    key,
		Object: value,
		// Value: value,
		Expiration: time.Hour,
		// TTL:   time.Hour,
	}); err != nil {
		return err
	}
	return nil
}

func NewInternalRedisImpl(rdb *redis.Ring) InternalRedis {
	return &InternalRedisImpl{
		rdb: rdb,
		cache: &cache.Codec{
			Redis: rdb,
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
			// LocalCache: cache.NewTinyLFU(1000, time.Minute),
		},
	}
}
