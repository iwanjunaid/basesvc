package author

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/iwanjunaid/basesvc/internal/respond"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

// Create handles HTTP POST requst for creating a new author
func Create(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var author *model.Author
		err := ctx.BodyParser(&author)

		if err != nil {
			return respond.Fail(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, err)
		}
		var ok bool
		if ok, err = isRequestValid(*author); !ok {
			return respond.Fail(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, err)
		}
		appController := rest.GetAppController()
		authorRes, err := appController.Author.InsertAuthor(ctx.Context(), author)
		if err != nil {
			return respond.Fail(ctx, http.StatusInternalServerError, http.StatusInternalServerError, err)
		}
		return respond.Success(ctx, http.StatusOK, authorRes)
	}
}

// isRequestValid validates Author struct
func isRequestValid(a model.Author) (bool, error) {
	var err error

	// validation using ozzo-validation module
	// more info: https://github.com/go-ozzo/ozzo-validation
	err = validation.ValidateStruct(&a,
		validation.Field(&a.ID, validation.Required),
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Email, validation.Required, is.Email))

	if err != nil {
		return false, err
	}

	err = validation.Validate(a.Email, validation.By(onlyGmail))
	if err != nil {
		return false, err
	}

	return true, nil
}

// onlyGmail validates email must gmail
func onlyGmail(value interface{}) error {
	s, _ := value.(string)

	re := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@gmail\.com`)

	if !re.MatchString(s) {
		return errors.New("email is not gmail")
	}

	return nil
}
