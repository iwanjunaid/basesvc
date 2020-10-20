package repository

import (
	"context"
	"github.com/iwanjunaid/basesvc/domain/model"
)

// ArticleRepository :
type ArticleRepository interface{
	FindAll(c context.Context) ([]domain.Article, error)
}