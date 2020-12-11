package registry

import (
	gp "github.com/iwanjunaid/basesvc/adapter/presenter"
	cr "github.com/iwanjunaid/basesvc/adapter/repository/cache"
	"github.com/iwanjunaid/basesvc/usecase/gravatar/interactor"
	"github.com/iwanjunaid/basesvc/usecase/gravatar/presenter"
	"github.com/iwanjunaid/basesvc/usecase/gravatar/repository"
)

func (r *registry) NewGravatarInteractor() interactor.GravatarInteractor {
	return interactor.NewGravatarInteractor(r.NewGravatarPresenter(),
		interactor.GravatarCacheRepository(r.NewGravatarCacheRepository()),
	)
}

func (r *registry) NewGravatarPresenter() presenter.GravatarPresenter {
	return gp.NewGravatarPresenter()
}

func (r *registry) NewGravatarCacheRepository() repository.GravatarCacheRepository {
	return cr.NewGravatarCacheRepository(r.rdb)
}
