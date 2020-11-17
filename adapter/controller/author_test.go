package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	// adapter "github.com/iwanjunaid/basesvc/shared/mock/adapter"

	"github.com/iwanjunaid/basesvc/shared/mock/interactor"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertAuthorController(t *testing.T) {
	Convey("Insert Author Controller", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		authorInteractor := interactor.NewMockAuthorInteractor(ctrl)
		Convey("Negative Scenarios", func() {
			Convey("Should return error", func() {
				//entAuthor := &model.Author{
				//	Name:      "123",
				//	Email:     "123",
				//	CreatedAt: time.Now().Unix(),
				//	UpdatedAt: time.Now().Unix(),
				//}

				authorInteractor.EXPECT().Create(context.Background(), nil).Return(nil, errors.New("error"))
				ctrl := NewAuthorController(authorInteractor)
				_, err := ctrl.InsertAuthor(context.Background(), nil)
				So(err, ShouldNotBeNil)
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
