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
	rd "github.com/iwanjunaid/basesvc/shared/mock/intern/redis"
	"github.com/iwanjunaid/basesvc/shared/mock/repository"

	in "github.com/iwanjunaid/basesvc/usecase/author/interactor"
	gi "github.com/iwanjunaid/basesvc/usecase/gravatar/interactor"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertAuthorController(t *testing.T) {
	Convey("Insert Author Controller", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repoAuthor := repository.NewMockAuthorSQLRepository(ctrl)
		repoCacheGravatar := repository.NewMockGravatarCacheRepository(ctrl)
		repoEventAuthor := repository.NewMockAuthorEventRepository(ctrl)
		Convey("Negative Scenarios", func() {
			Convey("Should return error", func() {
				repoAuthor.EXPECT().Create(context.Background(), nil).Return(nil, errors.New("error"))
				repoEventAuthor.EXPECT().Publish(context.Background(), nil, nil).Return(errors.New("error")).AnyTimes()
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorSQLRepository(repoAuthor), in.AuthorEventRepository(repoEventAuthor))
				gvCtrl := gi.NewGravatarInteractor(nil, gi.GravatarCacheRepository(repoCacheGravatar))
				svc := NewAuthorController(auCtrl, gvCtrl)
				res, err := svc.InsertAuthor(context.Background(), nil)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Should return error", func() {
				entAuthor := &model.Author{
					Name:      "123",
					Email:     "email2@gmail.com",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				}
				repoAuthor.EXPECT().Create(context.Background(), entAuthor).Return(entAuthor, nil)
				entByte, _ := json.Marshal(entAuthor)
				repoEventAuthor.EXPECT().Publish(context.Background(), nil, entByte).Return(nil)
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorSQLRepository(repoAuthor), in.AuthorEventRepository(repoEventAuthor))
				gvCtrl := gi.NewGravatarInteractor(nil, gi.GravatarCacheRepository(repoCacheGravatar))
				svc := NewAuthorController(auCtrl, gvCtrl)
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
		repoCacheAuthor := rd.NewMockInternalRedis(ctrl)
		repoCacheGravatar := repository.NewMockGravatarCacheRepository(ctrl)
		presenterAuthor := presenter.NewAuthorPresenter()
		Convey("Negative Scenarios", func() {
			Convey("Should return error", func() {
				var author []*model.Author
				repoCacheAuthor.EXPECT().Get(context.Background(), "all_authors", &author).Return(errors.New("error"))
				repoAuthor.EXPECT().FindAll(context.Background()).Return(nil, errors.New("error"))
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorSQLRepository(repoAuthor), in.AuthorCacheRepository(repoCacheAuthor))
				gvCtrl := gi.NewGravatarInteractor(nil, gi.GravatarCacheRepository(repoCacheGravatar))
				svc := NewAuthorController(auCtrl, gvCtrl)
				res, err := svc.GetAuthors(context.Background())
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Should return error", func() {
				var entAuthor, author []*model.Author
				entAuthor = append(entAuthor, &model.Author{
					Name:      "123",
					Email:     "email2@gmail.com",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				})

				entProfile := &model.GravatarProfiles{
					Entry: []model.Profile{
						{
							"120749118",
							"cb4c9309231b46cca2d6ee14303a7679",
							"cb4c9309231b46cca2d6ee14303a7679",
							"http://gravatar.com/yeninesilcik",
							"https://secure.gravatar.com/avatar/cb4c9309231b46cca2d6ee14303a7679",
							[]model.Photo{
								{
									"https://secure.gravatar.com/avatar/cb4c9309231b46cca2d6ee14303a7679",
									"thumbnail",
								},
							},
							[]string{},
							"yeninesilcik",
							[]string{},
						},
					},
				}
				repoCacheAuthor.EXPECT().Get(context.Background(), "all_authors", &author).Return(errors.New("error"))
				repoAuthor.EXPECT().FindAll(context.Background()).Return(entAuthor, nil)
				repoCacheAuthor.EXPECT().Create(context.Background(), "all_authors", entAuthor).Return(nil)
				repoCacheGravatar.EXPECT().Find(context.Background(), "0713fe419c9cdf793cd8c2d6e50ac07c21059d5364535387ae46f4003850294a").Return(nil, errors.New("error"))
				repoCacheGravatar.EXPECT().Create(context.Background(), "0713fe419c9cdf793cd8c2d6e50ac07c21059d5364535387ae46f4003850294a", entProfile).Return(nil)
				presenterAuthor.ResponseUsers(context.Background(), entAuthor)
				auCtrl := in.NewAuthorInteractor(presenterAuthor, in.AuthorSQLRepository(repoAuthor), in.AuthorCacheRepository(repoCacheAuthor))
				gvCtrl := gi.NewGravatarInteractor(nil, gi.GravatarCacheRepository(repoCacheGravatar))
				svc := NewAuthorController(auCtrl, gvCtrl)
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
		repoCacheGravatar := repository.NewMockGravatarCacheRepository(ctrl)
		Convey("Negative Scenarios", func() {
			Convey("Should return error", func() {
				repoAuthor.EXPECT().Create(context.Background(), nil).Return(errors.New("error"))
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorDocumentRepository(repoAuthor))
				gvCtrl := gi.NewGravatarInteractor(nil, gi.GravatarCacheRepository(repoCacheGravatar))
				svc := NewAuthorController(auCtrl, gvCtrl)
				res := svc.InsertDocument(context.Background(), nil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Should return error", func() {
				entAuthor := &model.Author{
					Name:      "123",
					Email:     "email2@gmail.com",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				}
				repoAuthor.EXPECT().Create(context.Background(), entAuthor).Return(nil)
				auCtrl := in.NewAuthorInteractor(nil, in.AuthorDocumentRepository(repoAuthor))
				gvCtrl := gi.NewGravatarInteractor(nil, gi.GravatarCacheRepository(repoCacheGravatar))
				svc := NewAuthorController(auCtrl, gvCtrl)
				res := svc.InsertDocument(context.Background(), entAuthor)
				So(res, ShouldBeNil)
			})
		})
	})
}
