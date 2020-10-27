package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi/middleware"
)

const ErrMaxStack = 5

type logrusRequestLoggerEntry struct {
	Logger log.FieldLogger
}

func (l *logrusRequestLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.WithFields(log.Fields{
		"resp_status": status, "resp_bytes_length": bytes,
		"resp_elapsed_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})
	l.Logger.Infoln("request complete")
}

func (l *logrusRequestLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(log.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})

}

func LogEntrySetFields(r *fiber.Ctx, fields map[string]interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*logrusRequestLoggerEntry); ok {
		entry.Logger = entry.Logger.WithFields(fields)
	}
}
