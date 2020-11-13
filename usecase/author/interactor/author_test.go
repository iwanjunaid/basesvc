package interactor

import (
	"context"
	"errors"
	"testing"

	mockRepo "github.com/iwanjunaid/basesvc/shared/mock/repository"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"
	. "github.com/smartystreets/goconvey/convey"
)

// func setupAuthorPositive() {
// 	entAuthor = &model.Author{
// 		ID:        1,
// 		Name:      "123",
// 		Email:     "123",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// }

func TestAuthor(t *testing.T) {
	Convey("Insert Author", t, func() {
		repoAuthor := &mockRepo.MockAuthorRepository{}
		Convey("Negative Scenarios", func() {
			Convey("Should return error ", func() {
				repoAuthor.On("Create", context.Background(), nil).Return(nil, errors.New("error"))
				uc := NewAuthorInteractor(nil, interactor.AuthorSQLRepository())
				// _, err := uc.Create(nil, nil)
				So(uc, ShouldBeNil)
			})
		})
		// Convey("Positive Scenarios", func() {
		// 	Convey("Insert Author", func() {
		// 		setupNpgCardPositive()
		// 		repoAuthor.On("Create", mock.Anything).Return(nil)
		// 		uc := NewAuthorInteractor(repoAuthor)
		// 		err := usecase.interactor.(entAuthor)
		// 		So(err, ShouldBeNil)
		// 	})
		// })
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
