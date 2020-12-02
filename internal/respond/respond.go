package respond

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/telemetry"

	validation "github.com/go-ozzo/ozzo-validation"
)

const ErrMaxStack = 5

type (
	Causer interface {
		Cause() error
	}
	Response struct {
		RequestId string      `json:"request_id"`
		Content   interface{} `json:"content"`
		Error     *Error      `json:"error"`
		Status    int         `json:"status"`
	}
	Error struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Reasons validation.Errors `json:"reasons"`
	}
)

func (err *Error) Error() string {
	return fmt.Sprintf("error with code: %d; message: %s", err.Code, err.Message)
}

func Success(c *fiber.Ctx, status int, content interface{}) error {
	txn := telemetry.GetTelemetry(c.Context())
	defer txn.End()

	return c.JSON(&Response{
		RequestId: c.Context().Value("requestid").(string),
		Status:    status,
		Content:   content,
	})
}

func Fail(c *fiber.Ctx, status, errorCode int, err error) error {
	var (
		message = err.Error()
		reason  = validation.Errors{}
	)
	txn := telemetry.GetTelemetry(c.Context())
	defer txn.End()
	txn.NoticeError(err)

	// if error masked, get detail!
	if ec, ok := err.(Causer); ok {
		err = ec.Cause()
	}
	if ev, ok2 := err.(validation.Errors); ok2 {
		message = "there`s some validation issues in request attributes"
		reason = ev
	}
	c.Status(status)
	return c.JSON(&Response{
		RequestId: c.Context().Value("requestid").(string),
		Status:    status,
		Error: &Error{
			Code:    errorCode,
			Message: message,
			Reasons: reason,
		},
	})

}
