package registry

import (
	"github.com/iwanjunaid/basesvc/adapter/controller"
	ap "github.com/iwanjunaid/basesvc/adapter/presenter"
	cr "github.com/iwanjunaid/basesvc/adapter/repository/cache"
	dr "github.com/iwanjunaid/basesvc/adapter/repository/document"
	er "github.com/iwanjunaid/basesvc/adapter/repository/event"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"
	"github.com/iwanjunaid/basesvc/usecase/author/presenter"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

func (r *registry) NewAuthorController() controller.AuthorController {
	return controller.NewAuthorController(r.NewAuthorInteractor())
}

func (r *registry) NewAuthorInteractor() interactor.AuthorInteractor {
	return interactor.NewAuthorInteractor(r.NewAuthorPresenter(),
		interactor.AuthorSQLRepository(r.NewAuthorSQLRepository()),
		interactor.AuthorDocumentRepository(r.NewAuthorDocumentRepository()),
		interactor.AuthorCacheRepository(r.NewAuthorCacheRepository()),
		interactor.AuthorEventRepository(r.NewEventRepository()))
}

func (r *registry) NewAuthorSQLRepository() repository.AuthorSQLRepository {
	return sr.NewAuthorRepository(r.db)
}

func (r *registry) NewAuthorDocumentRepository() repository.AuthorDocumentRepository {
	return dr.NewAuthorDocumentRepository(r.mdb)
}

func (r *registry) NewAuthorCacheRepository() repository.AuthorCacheRepository {
	return cr.NewAuthorCacheRepository()
}

func (r *registry) NewEventRepository() repository.AuthorEventRepository {
	return er.NewAuthorEventRepository(r.kP)
}

func (r *registry) NewAuthorPresenter() presenter.AuthorPresenter {
	return ap.NewAuthorPresenter()
}
