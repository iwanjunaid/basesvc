package registry

import (
	"github.com/iwanjunaid/basesvc/adapter/controller"
	ap "github.com/iwanjunaid/basesvc/adapter/presenter"
	dr "github.com/iwanjunaid/basesvc/adapter/repository/document"
	sr "github.com/iwanjunaid/basesvc/adapter/repository/sql"
	"github.com/iwanjunaid/basesvc/internal/kafka"
	rInternal "github.com/iwanjunaid/basesvc/internal/redis"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"
	"github.com/iwanjunaid/basesvc/usecase/author/presenter"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

func (r *registry) NewAuthorController() controller.AuthorController {
	return controller.NewAuthorController(
		r.NewAuthorInteractor(),
		r.NewGravatarInteractor(),
	)
}

func (r *registry) NewAuthorInteractor() interactor.AuthorInteractor {
	return interactor.NewAuthorInteractor(r.NewAuthorPresenter(),
		interactor.AuthorSQLRepository(r.NewAuthorSQLRepository()),
		interactor.AuthorDocumentRepository(r.NewAuthorDocumentRepository()),
		interactor.AuthorCacheRepository(r.NewAuthorCacheRepository()),
		interactor.AuthorEventRepository(r.NewEventRepository()),
	)
}

func (r *registry) NewAuthorSQLRepository() repository.AuthorSQLRepository {
	return sr.NewAuthorRepository(r.db)
}

func (r *registry) NewAuthorDocumentRepository() repository.AuthorDocumentRepository {
	return dr.NewAuthorDocumentRepository(r.mdb)
}

func (r *registry) NewAuthorCacheRepository() rInternal.InternalRedis {
	return rInternal.NewInternalRedisImpl(r.rdb)
}

func (r *registry) NewEventRepository() kafka.InternalKafka {
	return kafka.NewInternalKafkaImpl(r.kP)
}

func (r *registry) NewAuthorPresenter() presenter.AuthorPresenter {
	return ap.NewAuthorPresenter()
}
