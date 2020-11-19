package cache

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

type AuthorCacheRepositoryImpl struct {
	rdb   *redis.Ring
	cache *cache.Cache
}

func (ar *AuthorCacheRepositoryImpl) Find(ctx context.Context, key string) (author *model.Author, err error) {
	if err := ar.cache.Get(ctx, key, &author); err != nil {
		return nil, err
	}
	return author, nil
}

func (ar *AuthorCacheRepositoryImpl) FindAll(ctx context.Context, key string) (authors []*model.Author, err error) {
	if err := ar.cache.Get(ctx, key, &authors); err != nil {
		return nil, err
	}

	return authors, nil
}

func (ar *AuthorCacheRepositoryImpl) Create(ctx context.Context, key string, value interface{}) error {
	if err := ar.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time.Hour,
	}); err != nil {
		return err
	}
	return nil
}

func NewAuthorCacheRepository(rdb *redis.Ring) repository.AuthorCacheRepository {
	return &AuthorCacheRepositoryImpl{
		rdb: rdb,
		cache: cache.New(&cache.Options{
			Redis:      rdb,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		}),
	}
}
