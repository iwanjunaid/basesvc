package author

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

// Create handles HTTP POST requst for creating a new author
func Create(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var author model.Author
		err := ctx.BodyParser(&author)

		if err != nil {
			ctx.Response().SetStatusCode(http.StatusUnprocessableEntity)
			return err
		}

		var ok bool
		if ok, err = isRequestValid(author); !ok {
			return err
		}

		appController := rest.GetAppController()
		return appController.Author.CreateAuthor(ctx, &author)
	}
}

// isRequestValid validates Author struct
func isRequestValid(a model.Author) (bool, error) {
	var err error
	err = validation.ValidateStruct(&a,
		validation.Field(&a.ID, validation.Required),
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Email, validation.Required, is.Email))

	if err != nil {
		return false, err
	}

	return true, nil
}
