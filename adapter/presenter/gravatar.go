package presenter

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/gravatar/presenter"
)

type GravatarPresenterImpl struct{}

func NewGravatarPresenter() presenter.GravatarPresenter {
	return &GravatarPresenterImpl{}
}

func (presenter *GravatarPresenterImpl) ResponseGravatar(ctx context.Context, gravatar *model.GravatarProfiles) (*model.GravatarProfiles, error) {
	return gravatar, nil
}
