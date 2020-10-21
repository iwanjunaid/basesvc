package presenter

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

// ArticlePresenter :
type ArticlePresenter interface {
	ResponseArticle(c context.Context) []*model.Article
}