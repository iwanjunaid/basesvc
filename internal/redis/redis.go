package redis

import (
	"context"
	"time"

	"github.com/iwanjunaid/basesvc/internal/telemetry"

	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
)

type InternalRedisImpl struct {
	rdb   *redis.Ring
	cache *cache.Codec
}

type InternalRedis interface {
	Get(ctx context.Context, key string, obj interface{}) error
	Create(ctx context.Context, key string, value interface{}) error
}

func (ar *InternalRedisImpl) nrRedisHook(c context.Context) error {
	redis := telemetry.StartRedisSegment(c, ar.rdb)
	ar.rdb = redis
	ar.cache.Redis = redis
	return nil
}

func (ar *InternalRedisImpl) Get(ctx context.Context, key string, obj interface{}) error {
	ar.nrRedisHook(ctx)
	if err := ar.cache.GetContext(ctx, key, obj); err != nil {
		return err
	}
	return nil
}

func (ar *InternalRedisImpl) Create(ctx context.Context, key string, value interface{}) error {
	ar.nrRedisHook(ctx)
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
		},
	}
}
