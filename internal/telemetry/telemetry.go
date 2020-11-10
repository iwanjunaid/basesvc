package telemetry

import (
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

type PathFn func(r *http.Request) string

func Middleware(nra newrelic.Application, fn PathFn) func(next http.Handler) http.Handler {
	if fn == nil {
		fn = func(r *http.Request) string {
			return r.Method + " " + r.URL.Path
		}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if nra != nil {
				txn := nra.StartTransaction(fn(r), w, r)
				defer func(t newrelic.Transaction) {
					_ = t.End()
				}(txn)
				r = r.WithContext(newrelic.NewContext(r.Context(), txn))
			}
			next.ServeHTTP(w, r)
		})
	}
}
