package logger

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	logger "github.com/sirupsen/logrus"
)

const ErrMaxStack = 5

var (
	// LogEntryCtxKey is the context.Context key to store the request log entry.
	LogEntryCtxKey = "LogEntry"

	// DefaultLogger is called by the Logger middleware handler to log each request.
	// Its made a package-level variable so that it can be reconfigured for custom
	// logging configurations.
)

type logrusRequestLogger struct {
	Logger *logger.Logger
}

// Logger is a middleware that logs the start and end of each request, along
// with some useful data about what was requested, what the response status was,
// and how long it took to return. When standard output is a TTY, Logger will
// print in color, otherwise it will print in black and white. Logger prints a
// request ID if one is provided.
//
// Alternatively, look at https://github.com/pressly/lg and the `lg.RequestLogger`
// middleware pkg.

func requestLogger(f LogFormatter) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		entry := f.NewLogEntry(ctx)
		ctx.Locals(LogEntryCtxKey, entry)

		t1 := time.Now()
		defer func() {
			entry.Write(ctx.Response().StatusCode(), time.Since(t1))
		}()

		return ctx.Next()
	}
}

// LogFormatter initiates the beginning of a new LogEntry per request.
// See DefaultLogFormatter for an example implementation.
type LogFormatter interface {
	NewLogEntry(r *fiber.Ctx) LogEntry
}

// LogEntry records the final log when a request completes.
// See defaultLogEntry for an example implementation.
type LogEntry interface {
	Write(status int, elapsed time.Duration)
	Panic(v interface{}, stack []byte)
}

func RequestLogger(logParam *logger.Logger) fiber.Handler {
	return requestLogger(&logrusRequestLogger{Logger: logParam})
}

// GetLogEntry returns the in-context LogEntry for a request.
func GetLogEntry(r *http.Request) LogEntry {
	entry, _ := r.Context().Value(LogEntryCtxKey).(LogEntry)
	return entry
}

// WithLogEntry sets the in-context LogEntry for a request.
func WithLogEntry(r *http.Request, entry LogEntry) *http.Request {
	r = r.WithContext(context.WithValue(r.Context(), LogEntryCtxKey, entry))
	return r
}

// LoggerInterface accepts printing to stdlib logger or compatible logger.
type LoggerInterface interface {
	Print(v ...interface{})
}

// DefaultLogFormatter is a simple logger that implements a LogFormatter.
type DefaultLogFormatter struct {
	Logger  LoggerInterface
	NoColor bool
}

func (l *logrusRequestLogger) NewLogEntry(r *fiber.Ctx) LogEntry {
	entry := &logrusRequestLoggerEntry{Logger: logger.NewEntry(l.Logger)}
	logFields := logger.Fields{}

	logFields["ts"] = time.Now().UTC().Format(time.RFC1123)

	if reqID := r.Get("X-Request-ID"); reqID != "" {
		logFields["req_id"] = reqID
	}
	scheme := "http"
	if r.Protocol() != "" {
		scheme = "https"
	}
	logFields["http_scheme"] = scheme
	//logFields["http_method"] = r.Method
	logFields["remote_addr"] = r.IP()
	logFields["user_agent"] = r.Get("Origin")
	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Hostname(), r.BaseURL())

	entry.Logger = entry.Logger.WithFields(logFields)

	return entry
}

type logrusRequestLoggerEntry struct {
	Logger logger.FieldLogger
}

func (l *logrusRequestLoggerEntry) Write(status int, elapsed time.Duration) {
	l.Logger = l.Logger.WithFields(logger.Fields{
		"resp_status":     status,
		"resp_elapsed_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})
	l.Logger.Infoln("request complete")
}

func (l *logrusRequestLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logger.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}

func LogEntrySetFields(r *fiber.Ctx, fields map[string]interface{}) {
	if entry, ok := r.Context().Value(LogEntryCtxKey).(*logrusRequestLoggerEntry); ok {
		entry.Logger = entry.Logger.WithFields(fields)
	}
}
