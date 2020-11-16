package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	// adapter "github.com/iwanjunaid/basesvc/shared/mock/adapter"

	repository "github.com/iwanjunaid/basesvc/shared/mock/repository"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertAuthorController(t *testing.T) {
	Convey("Insert Author Controller", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repoAuthor := repository.NewMockAuthorSQLRepository(ctrl)
		// adapter := adapter.NewMockAuthorController(ctrl)

		Convey("Negative Scenarios", func() {
			Convey("Should return error", func() {
				repoAuthor.EXPECT().Create(context.Background(), nil).Return(nil, errors.New("error"))
				auCtrl := interactor.NewAuthorInteractor(nil, interactor.AuthorSQLRepository(repoAuthor))
				svc := NewAuthorController(auCtrl)
				res := svc.InsertAuthor(nil)
				So(res, ShouldNotBeNil)
			})
		})
		// Convey("Positive Scenarios", func() {
		// 	Convey("Should return error", func() {
		// 		entAuthor := &model.Author{
		// 			Name:      "123",
		// 			Email:     "123",
		// 			CreatedAt: time.Now(),
		// 			UpdatedAt: time.Now(),
		// 		}
		// 		repoAuthor.EXPECT().Create(context.Background(), entAuthor).Return(entAuthor, errors.New("error"))
		// 		auCtrl := interactor.NewAuthorInteractor(nil, interactor.AuthorSQLRepository(repoAuthor))
		// 		svc := NewAuthorController(auCtrl)
		// 		res := svc.InsertAuthor(nil)
		// 		So(res, ShouldBeNil)
		// 	})
		// })
	})
}
