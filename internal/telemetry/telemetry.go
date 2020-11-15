package telemetry

import (
	"context"
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

type PathFn func(r *http.Request) string

const telemetryTxnCtxKey = "telemetry"

func Middleware(nra newrelic.Application, fn PathFn) func(next http.Handler) http.Handler {
	if fn == nil {
		fn = func(r *http.Request) string {
			return r.Method + " " + r.URL.Path
		}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if nra != nil {
				txn := nra.StartTransaction(fn(r), w, r)
				defer func(t newrelic.Transaction) {
					_ = t.End()
				}(txn)
				r = r.WithContext(newrelic.NewContext(r.Context(), txn))
				ctx = context.WithValue(r.Context(), telemetryTxnCtxKey, txn)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
