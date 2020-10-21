package registry

import (
	"github.com/iwanjunaid/basesvc/adapter/controller"
	ap "github.com/iwanjunaid/basesvc/adapter/presenter"
	ar "github.com/iwanjunaid/basesvc/adapter/repository"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"
	"github.com/iwanjunaid/basesvc/usecase/author/presenter"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

func (r *registry) NewAuthorController() controller.AuthorController {
	return controller.NewAuthorController(r.NewAuthorInteractor())
}

func (r *registry) NewAuthorInteractor() interactor.AuthorInteractor {
	return interactor.NewAuthorInteractor(r.NewAuthorRepository(), r.NewAuthorPresenter())
}

func (r *registry) NewAuthorRepository() repository.AuthorRepository {
	return ar.NewAuthorRepository(r.db)
}

func (r *registry) NewAuthorPresenter() presenter.AuthorPresenter {
	return ap.NewAuthorPresenter()
}
