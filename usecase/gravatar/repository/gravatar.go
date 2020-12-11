package repository

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

type GravatarRepository interface {
	GetProfile() (*model.GravatarProfiles, error)
	URL() (string, error)
	JSONURL() (string, error)
	AvatarURL() (string, error)
}

type GravatarCacheRepository interface {
	Find(ctx context.Context, key string) (*model.GravatarProfiles, error)
	Create(ctx context.Context, key string, value interface{}) error
}
