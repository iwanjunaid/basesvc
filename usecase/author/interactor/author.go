package interactor

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/iwanjunaid/basesvc/adapter/repository/api"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/presenter"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

type AuthorInteractor interface {
	Get(ctx context.Context, key string, id string) (*model.Author, error)
	GetAll(ctx context.Context, key string) ([]*model.Author, error)
	CreateDocument(ctx context.Context, author *model.Author) error
	Create(ctx context.Context, author *model.Author) (*model.Author, error)
}

type AuthorInteractorImpl struct {
	AuthorSQLRepository           repository.AuthorSQLRepository
	AuthorDocumentRepository      repository.AuthorDocumentRepository
	AuthorCacheRepository         repository.AuthorCacheRepository
	AuthorEventRepository         repository.AuthorEventRepository
	AuthorGravatarRepository      repository.AuthorGravatarRepository
	AuthorGravatarCacheRepository repository.AuthorGravatarCacheRepository
	AuthorPresenter               presenter.AuthorPresenter
}

type Option func(impl *AuthorInteractorImpl)

func NewAuthorInteractor(p presenter.AuthorPresenter, option ...Option) AuthorInteractor {
	interactor := &AuthorInteractorImpl{
		AuthorPresenter: p,
	}
	for _, opt := range option {
		opt(interactor)
	}
	return interactor
}

func AuthorSQLRepository(sql repository.AuthorSQLRepository) Option {
	return func(impl *AuthorInteractorImpl) {
		impl.AuthorSQLRepository = sql
	}
}

func AuthorDocumentRepository(document repository.AuthorDocumentRepository) Option {
	return func(impl *AuthorInteractorImpl) {
		impl.AuthorDocumentRepository = document
	}
}

func AuthorCacheRepository(cache repository.AuthorCacheRepository) Option {
	return func(impl *AuthorInteractorImpl) {
		impl.AuthorCacheRepository = cache
	}
}

func AuthorGravatarCacheRepository(cache repository.AuthorGravatarCacheRepository) Option {
	return func(impl *AuthorInteractorImpl) {
		impl.AuthorGravatarCacheRepository = cache
	}
}

func AuthorEventRepository(event repository.AuthorEventRepository) Option {
	return func(impl *AuthorInteractorImpl) {
		impl.AuthorEventRepository = event
	}
}

func (ai *AuthorInteractorImpl) setAvatar(ctx context.Context, author *model.Author) (*model.Author, error) {
	// Set Gravatar Profile
	var avatar string

	profile, err := ai.getAvatar(ctx, author)

	if profile != nil && len(profile.Entry) > 0 {
		avatar = profile.Entry[0].ThumbnailUrl
	}

	if err != nil {
		return author, err
	}
	author.Avatar = avatar

	return author, nil
}

func (ai *AuthorInteractorImpl) getAvatar(ctx context.Context, author *model.Author) (profile *model.GravatarProfiles, err error) {
	// Get Gravatar Profile
	gravatar := api.NewAuthorGravatar(ctx, author.Email)

	profile = &model.GravatarProfiles{}

	// Key
	h := sha256.New()
	h.Write([]byte(author.Email))
	key := fmt.Sprintf("%x", h.Sum(nil))

	// Get cache from redis
	profile, err = ai.AuthorGravatarCacheRepository.Find(ctx, key)

	if profile != nil {
		return profile, nil
	}

	profile, err = gravatar.GetProfile()
	if err != nil {
		return nil, err
	}

	// Set value for the key
	if err = ai.AuthorGravatarCacheRepository.Create(ctx, key, profile); err != nil {
		return profile, err
	}

	return profile, nil
}

func (ai *AuthorInteractorImpl) Get(ctx context.Context, key string, id string) (author *model.Author, err error) {
	defer func() {
		author, _ = ai.setAvatar(ctx, author)
	}()

	// Get value from redis based on the key
	author, err = ai.AuthorCacheRepository.Find(ctx, key)

	if author != nil {
		return ai.AuthorPresenter.ResponseUser(ctx, author)
	}

	// If not available then get from sql repo
	author, err = ai.AuthorSQLRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	// Set value for the key
	if err = ai.AuthorCacheRepository.Create(ctx, key, author); err != nil {
		return author, err
	}

	if err != nil {
		return author, err
	}

	return ai.AuthorPresenter.ResponseUser(ctx, author)
}

func (ai *AuthorInteractorImpl) GetAll(ctx context.Context, key string) (authors []*model.Author, err error) {
	defer func() {
		for _, author := range authors {
			author, _ = ai.setAvatar(ctx, author)
		}
	}()

	// Get value from redis based on the key
	authors, err = ai.AuthorCacheRepository.FindAll(ctx, key)

	if authors != nil {
		return ai.AuthorPresenter.ResponseUsers(ctx, authors)
	}

	// If not available then get from sql repo
	authors, err = ai.AuthorSQLRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// Set value for the key
	if err = ai.AuthorCacheRepository.Create(ctx, key, authors); err != nil {
		return authors, err
	}
	return ai.AuthorPresenter.ResponseUsers(ctx, authors)
}

func (ai *AuthorInteractorImpl) CreateDocument(ctx context.Context, author *model.Author) error {
	err := ai.AuthorDocumentRepository.Create(ctx, author)
	if err != nil {
		return err
	}
	return nil
}

func (ai *AuthorInteractorImpl) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	var (
		err error
	)
	author, err = ai.AuthorSQLRepository.Create(ctx, author)
	if err != nil {
		return author, err
	}
	dataByt, err := json.Marshal(author)
	if err != nil {
		return author, err
	}

	if err := ai.AuthorEventRepository.Publish(ctx, nil, dataByt); err != nil {
		return author, err
	}
	return author, err
}
