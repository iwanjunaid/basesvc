package telemetry

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v7"
	"github.com/iwanjunaid/basesvc/config"

	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/gofiber/fiber/v2"

	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

type PathFn func(r *http.Request) string

type endable interface {
	End() error
}

const telemetryTxnCtxKey = "newRelicTransaction"

func NewrelicMiddleware(nra *newrelic.Application, fn PathFn) fiber.Handler {
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
				if nra != nil {
					txn := nra.StartTransaction(fn(r))
					// defer txn.End()

					// This marks the transaction as a web transactions and collects details on
					// the request attributes
					txn.SetWebRequestHTTP(r)
					// This collects details on response code and headers. Use the returned
					// Writer from here on.
					w = txn.SetWebResponse(w)

					c.Locals(telemetryTxnCtxKey, txn)
				}

				next.ServeHTTP(w, r)
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

func GetTelemetry(c context.Context) *newrelic.Transaction {
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
	// defer s.End()
	return
}

// StopDataSegment stops newrelic data store segment
func StopDataSegment(s endable) (err error) {
	if s != nil {
		err = s.End()
	}
	return err
}

// StartRedisSegment starts newrelic datastore redis for newrelic transaction
func StartRedisSegment(c context.Context, rdb *redis.Ring) *redis.Ring {
	txn := GetTelemetry(c)
	ctx := newrelic.NewContext(rdb.Context(), txn)

	return rdb.WithContext(ctx)
}
