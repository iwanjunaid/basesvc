package interactor

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/iwanjunaid/basesvc/adapter/presenter"
	"github.com/iwanjunaid/basesvc/domain/model"
	repository "github.com/iwanjunaid/basesvc/shared/mock/repository"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSQLAuthor(t *testing.T) {
	Convey("Insert Author", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repoAuthor := repository.NewMockAuthorSQLRepository(ctrl)
		repoEventAuthor := repository.NewMockAuthorEventRepository(ctrl)
		Convey("Negative Scenarios", func() {
			Convey("Should return error ", func() {
				repoAuthor.EXPECT().Create(context.Background(), nil).Return(nil, errors.New("error"))
				repoEventAuthor.EXPECT().Publish(context.Background(), nil, nil).Return(errors.New("error")).AnyTimes()
				uc := NewAuthorInteractor(nil, AuthorSQLRepository(repoAuthor), AuthorEventRepository(repoEventAuthor))
				_, err := uc.Create(context.Background(), nil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Insert Author", func() {
				entAuthor := &model.Author{
					Name:      "123",
					Email:     "123",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				}
				entByte, _ := json.Marshal(entAuthor)
				repoAuthor.EXPECT().Create(context.Background(), entAuthor).Return(entAuthor, nil)
				repoEventAuthor.EXPECT().Publish(context.Background(), nil, entByte).Return(nil)
				uc := NewAuthorInteractor(nil, AuthorSQLRepository(repoAuthor), AuthorEventRepository(repoEventAuthor))
				res, _ := uc.Create(context.Background(), entAuthor)
				So(res, ShouldEqual, entAuthor)
			})
		})
	})
}

func TestGetSQLAuthor(t *testing.T) {
	Convey("Get All Author", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repoAuthor := repository.NewMockAuthorSQLRepository(ctrl)
		presenterAuthor := presenter.NewAuthorPresenter()
		Convey("Negative Scenarios", func() {
			Convey("Should return error ", func() {
				repoAuthor.EXPECT().FindAll(context.Background()).Return(nil, errors.New("error"))
				uc := NewAuthorInteractor(nil, AuthorSQLRepository(repoAuthor))
				_, err := uc.GetAll(context.Background())
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Get All Author", func() {
				var entAuthor []*model.Author
				entAuthor = append(entAuthor, &model.Author{
					Name:      "123",
					Email:     "123",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				})
				repoAuthor.EXPECT().FindAll(context.Background()).Return(entAuthor, nil)
				presenterAuthor.ResponseUsers(context.Background(), entAuthor)
				uc := NewAuthorInteractor(presenterAuthor, AuthorSQLRepository(repoAuthor))
				res, _ := uc.GetAll(context.Background())
				So(res, ShouldNotBeNil)
			})
		})
	})
}

func TestMongoAuthor(t *testing.T) {
	Convey("Insert Author", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repoAuthor := repository.NewMockAuthorDocumentRepository(ctrl)
		Convey("Negative Scenarios", func() {
			Convey("Should return error ", func() {
				repoAuthor.EXPECT().Create(context.Background(), nil).Return(errors.New("error"))
				ucDoc := NewAuthorInteractor(nil, AuthorDocumentRepository(repoAuthor))
				err := ucDoc.CreateDocument(context.Background(), nil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Insert Author", func() {
				entAuthor := &model.Author{
					Name:      "123",
					Email:     "123",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				}
				repoAuthor.EXPECT().Create(context.Background(), entAuthor).Return(nil)
				ucDoc := NewAuthorInteractor(nil, AuthorDocumentRepository(repoAuthor))
				err := ucDoc.CreateDocument(context.Background(), entAuthor)
				So(err, ShouldBeNil)
			})
		})
	})
}
