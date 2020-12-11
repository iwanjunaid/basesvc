package cache

import (
	"context"
	"time"

	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/internal/telemetry"
	"github.com/iwanjunaid/basesvc/usecase/gravatar/repository"
	"github.com/vmihailenco/msgpack/v4"
)

type GravatarCacheRepositoryImpl struct {
	rdb   *redis.Ring
	cache *cache.Codec
}

func (gc *GravatarCacheRepositoryImpl) nrredisHook(c context.Context) error {
	redis := telemetry.StartRedisSegment(c, gc.rdb)
	gc.rdb = redis
	gc.cache.Redis = redis
	return nil
}

func (gc *GravatarCacheRepositoryImpl) Find(ctx context.Context, key string) (profile *model.GravatarProfiles, err error) {
	gc.nrredisHook(ctx)
	if err := gc.cache.Get(key, &profile); err != nil {
		return nil, err
	}
	return profile, nil
}

func (gc *GravatarCacheRepositoryImpl) Create(ctx context.Context, key string, value interface{}) error {
	gc.nrredisHook(ctx)
	if err := gc.cache.Set(&cache.Item{
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

func NewGravatarCacheRepository(rdb *redis.Ring) repository.GravatarCacheRepository {
	return &GravatarCacheRepositoryImpl{
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
