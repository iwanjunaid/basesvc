package interactor

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/iwanjunaid/basesvc/adapter/repository/api"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/gravatar/presenter"
	"github.com/iwanjunaid/basesvc/usecase/gravatar/repository"
)

type GravatarInteractor interface {
	Get(ctx context.Context, email string) (*model.GravatarProfiles, error)
}

type GravatarInteractorImpl struct {
	GravatarRepository      repository.GravatarRepository
	GravatarCacheRepository repository.GravatarCacheRepository
	GravatarPresenter       presenter.GravatarPresenter
}

type Option func(impl *GravatarInteractorImpl)

func NewGravatarInteractor(p presenter.GravatarPresenter, option ...Option) GravatarInteractor {
	interactor := &GravatarInteractorImpl{
		GravatarPresenter: p,
	}
	for _, opt := range option {
		opt(interactor)
	}
	return interactor
}

func GravatarCacheRepository(cache repository.GravatarCacheRepository) Option {
	return func(impl *GravatarInteractorImpl) {
		impl.GravatarCacheRepository = cache
	}
}

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (gv *GravatarInteractorImpl) Get(ctx context.Context, email string) (profile *model.GravatarProfiles, err error) {
	// Get Gravatar Profile
	gravatar := api.NewGravatarRepository(ctx, email)

	// Key
	key := hash(email)

	// Get cache from redis
	profile, err = gv.GravatarCacheRepository.Find(ctx, key)

	if profile != nil {
		return profile, nil
	}

	profile, err = gravatar.GetProfile()
	if err != nil {
		return nil, err
	}

	// Set value for the key
	if err = gv.GravatarCacheRepository.Create(ctx, key, profile); err != nil {
		return profile, err
	}

	return profile, nil
}
