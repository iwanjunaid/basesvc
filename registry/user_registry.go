package registry

import (
	"github.com/iwanjunaid/basesvc/interface/controller"
	ip "github.com/iwanjunaid/basesvc/interface/presenter"
	ir "github.com/iwanjunaid/basesvc/interface/repository"
	"github.com/iwanjunaid/basesvc/usecase/interactor"
	up "github.com/iwanjunaid/basesvc/usecase/presenter"
	ur "github.com/iwanjunaid/basesvc/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
