package interactor

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/iwanjunaid/basesvc/domain/model"
	repository "github.com/iwanjunaid/basesvc/shared/mock/repository/author"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"

	. "github.com/smartystreets/goconvey/convey"
)

//func setupAuthorPositive() {
//	entAuthor := &model.Author{
//		Name:      "123",
//		Email:     "123",
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//	}
//}

func TestAuthor(t *testing.T) {
	Convey("Insert Author", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repoAuthor := repository.NewMockAuthorSQLRepository(ctrl)
		Convey("Negative Scenarios", func() {
			Convey("Should return error ", func() {
				repoAuthor.EXPECT().Create(context.Background(), nil).Return(nil, errors.New("error"))
				uc := NewAuthorInteractor(nil, AuthorSQLRepository(repoAuthor))
				_, err := uc.Create(context.Background(), nil)
				So(err, ShouldBeError)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Insert Author", func() {
				entAuthor := &model.Author{
					Name:      "123",
					Email:     "123",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				repoAuthor.EXPECT().Create(context.Background(), entAuthor).Return(entAuthor, nil)
				//repoAuthor.On("Create", mock.Anything).Return(nil)
				uc := NewAuthorInteractor(nil, AuthorSQLRepository(repoAuthor))
				res, _ := uc.Create(context.Background(), entAuthor)
				So(res, ShouldEqual, entAuthor)
			})
		})
	})
}

// func TestAuthor(t *testing.T) {
// 	Convey("Author Test", t, func() {
// 		Convey("corner cases / negative scenarios", func() {
// 			Convey("case 1", func() {
// 				So("done", ShouldEqual, "done")
// 			})
// 		})
// 	})
// }
