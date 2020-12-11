package cache

import (
	"context"
	"time"

	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/internal/telemetry"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
	"github.com/vmihailenco/msgpack/v4"
)

type AuthorCacheRepositoryImpl struct {
	rdb   *redis.Ring
	cache *cache.Codec
}

func (ar *AuthorCacheRepositoryImpl) nrredisHook(c context.Context) error {
	redis := telemetry.StartRedisSegment(c, ar.rdb)
	ar.rdb = redis
	ar.cache.Redis = redis
	return nil
}

func (ar *AuthorCacheRepositoryImpl) Find(ctx context.Context, key string) (author *model.Author, err error) {
	ar.nrredisHook(ctx)
	if err := ar.cache.Get(key, &author); err != nil {
		return nil, err
	}
	return author, nil
}

func (ar *AuthorCacheRepositoryImpl) FindAll(ctx context.Context, key string) (authors []*model.Author, err error) {
	ar.nrredisHook(ctx)
	if err := ar.cache.Get(key, &authors); err != nil {
		return nil, err
	}

	return authors, nil
}

func (ar *AuthorCacheRepositoryImpl) Create(ctx context.Context, key string, value interface{}) error {
	ar.nrredisHook(ctx)
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

func NewAuthorCacheRepository(rdb *redis.Ring) repository.AuthorCacheRepository {
	return &AuthorCacheRepositoryImpl{
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
