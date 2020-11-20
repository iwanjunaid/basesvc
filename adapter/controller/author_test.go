package controller

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	// adapter "github.com/iwanjunaid/basesvc/shared/mock/adapter"

	"github.com/iwanjunaid/basesvc/adapter/presenter"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/shared/mock/repository"
	in "github.com/iwanjunaid/basesvc/usecase/author/interactor"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertAuthorController(t *testing.T) {
	Convey("Insert Author Controller", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repoAuthor := repository.NewMockAuthorSQLRepository(ctrl)
		repoEventAuthor := repository.NewMockAuthorEventRepository(ctrl)
		Convey("Negative Scenarios", func() {
			Convey("Should return error", func() {
				repoAuthor.EXPECT().Create(context.Background(), nil).Return(nil, errors.New("error"))
				repoEventAuthor.EXPECT().Publish(context.Background(), nil, nil).Return(errors.New("error")).AnyTimes()
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorSQLRepository(repoAuthor), in.AuthorEventRepository(repoEventAuthor))
				svc := NewAuthorController(auCtrl)
				res, err := svc.InsertAuthor(context.Background(), nil)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Should return error", func() {
				entAuthor := &model.Author{
					Name:      "123",
					Email:     "123",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				}
				repoAuthor.EXPECT().Create(context.Background(), entAuthor).Return(entAuthor, nil)
				entByte, _ := json.Marshal(entAuthor)
				repoEventAuthor.EXPECT().Publish(context.Background(), nil, entByte).Return(nil)
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorSQLRepository(repoAuthor), in.AuthorEventRepository(repoEventAuthor))
				svc := NewAuthorController(auCtrl)
				res, err := svc.InsertAuthor(context.Background(), entAuthor)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
	})
}

func TestGetAllAuthorController(t *testing.T) {
	Convey("Get All Author Controller", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repoAuthor := repository.NewMockAuthorSQLRepository(ctrl)
		presenterAuthor := presenter.NewAuthorPresenter()
		Convey("Negative Scenarios", func() {
			Convey("Should return error", func() {
				repoAuthor.EXPECT().FindAll(context.Background()).Return(nil, errors.New("error"))
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorSQLRepository(repoAuthor))
				svc := NewAuthorController(auCtrl)
				res, err := svc.GetAuthors(context.Background())
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Should return error", func() {
				var entAuthor []*model.Author
				entAuthor = append(entAuthor, &model.Author{
					Name:      "123",
					Email:     "123",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				})
				repoAuthor.EXPECT().FindAll(context.Background()).Return(entAuthor, nil)
				presenterAuthor.ResponseUsers(context.Background(), entAuthor)
				auCtrl := in.NewAuthorInteractor(presenterAuthor, in.AuthorSQLRepository(repoAuthor))
				svc := NewAuthorController(auCtrl)
				res, err := svc.GetAuthors(context.Background())
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
	})
}

func TestInsertAuthorDocumentController(t *testing.T) {
	Convey("Insert Author Document Controller", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repoAuthor := repository.NewMockAuthorDocumentRepository(ctrl)
		Convey("Negative Scenarios", func() {
			Convey("Should return error", func() {
				repoAuthor.EXPECT().Create(context.Background(), nil).Return(errors.New("error"))
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorDocumentRepository(repoAuthor))
				svc := NewAuthorController(auCtrl)
				res := svc.InsertDocument(context.Background(), nil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Should return error", func() {
				entAuthor := &model.Author{
					Name:      "123",
					Email:     "123",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				}
				repoAuthor.EXPECT().Create(context.Background(), entAuthor).Return(nil)
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorDocumentRepository(repoAuthor))
				svc := NewAuthorController(auCtrl)
				res := svc.InsertDocument(context.Background(), entAuthor)
				So(res, ShouldBeNil)
			})
		})
	})
}
