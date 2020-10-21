package registry

import (
	"database/sql"

	"github.com/iwanjunaid/basesvc/adapter/controller"
)

type registry struct {
	db *sql.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *sql.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Author: r.NewAuthorController(),
	}
}
