package telemetry

import (
	"context"
	"net/http"

	"github.com/iwanjunaid/basesvc/config"

	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/gofiber/fiber/v2"

	newrelic "github.com/newrelic/go-agent"
)

type PathFn func(r *http.Request) string

type endable interface {
	End() error
}

const telemetryTxnCtxKey = "telemetry"

func NewrelicMiddleware(nra newrelic.Application, fn PathFn) fiber.Handler {
	if fn == nil {
		fn = func(r *http.Request) string {
			return r.Method + " " + r.URL.Path
		}
	}
	return func(c *fiber.Ctx) error {
		var next bool
		nextHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) { next = true })
		_ = HTTPHandler(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var ctx = r.Context()
				if nra != nil {
					txn := nra.StartTransaction(fn(r), w, r)
					defer txn.End()
					//defer func(t newrelic.Transaction) {
					//	_ = t.End()
					//}(txn)
					r = r.WithContext(newrelic.NewContext(r.Context(), txn))
					c.Locals("newRelicTransaction", txn)
				}
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		}(nextHandler))(c)
		if next {
			return c.Next()
		}
		return nil
	}
}

// HTTPHandler wraps net/http handler to fiber handler
func HTTPHandler(h http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler := fasthttpadaptor.NewFastHTTPHandler(h)
		handler(c.Context())
		return nil
	}
}

func GetTelemetry(c context.Context) newrelic.Transaction {
	return newrelic.FromContext(c)
}

// StartDataSegment starts newrelic data store segment for newrelic transaction
func StartDataSegment(c context.Context, payload map[string]interface{}) (s *newrelic.DatastoreSegment) {
	nrt := GetTelemetry(c)
	if nrt == nil {
		return
	}
	s = &newrelic.DatastoreSegment{
		Product:            newrelic.DatastorePostgres,
		Collection:         payload["collection"].(string),
		Operation:          payload["operation"].(string),
		ParameterizedQuery: payload["query"].(string),
		QueryParameters:    payload["query_params"].(map[string]interface{}),
		Host:               config.GetString("database.postgres.host"),
		DatabaseName:       config.GetString("database.postgres.db"),
	}

	s.StartTime = nrt.StartSegmentNow()
	return
}

// StopDataSegment stops newrelic data store segment
func StopDataSegment(s endable) {
	if s != nil {
		_ = s.End()
	}
}
