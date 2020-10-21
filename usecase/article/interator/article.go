package interator

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/article/repository"
	"github.com/iwanjunaid/basesvc/usecase/article/presenter"
)


// ArticleInteractor : 
type ArticleInteractor interface {
	GetAll(c context.Context)
	GetByID(c context.Context, id uint) (model.Article, error)
	GetByTitle(c context.Context, title string) (model.Article, error)
}

// ArticleInteractorImpl :
type ArticleInteractorImpl struct {
	ArticleRepository repository.ArticleRepository
	ArticlePresenter repository.ArticlePresenter
}

func (ar *ArticleInteractorImpl) GetAll(c context.Context) (res []*model.Article, err error) {
	article, err := ar.ArticleRepository.FindAll(c)
	if err != nil {
		return 
	}

	return artc

}