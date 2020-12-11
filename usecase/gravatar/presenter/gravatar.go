package presenter

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

// GravatarPresenter :
type GravatarPresenter interface {
	ResponseGravatar(context.Context, *model.GravatarProfiles) (*model.GravatarProfiles, error)
}
