package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"

	"github.com/arturyumaev/file-processing/pkg/logger"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode  int
	wroteHeader bool
	body        []byte
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Write(bytes []byte) (int, error) {
	rw.body = bytes
	return rw.ResponseWriter.Write(bytes)
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func Logger(next http.Handler) http.Handler {
	logger := logger.Get()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		path := r.URL.Path
		rawQuery := r.URL.RawQuery
		if rawQuery != "" {
			path = path + "?" + rawQuery
		}

		wrappedRW := wrapResponseWriter(w)
		next.ServeHTTP(wrappedRW, r)

		requestId := r.Context().Value(ContextKeyRequestID).(string)
		timeAfter := time.Now()
		latency := timeAfter.Sub(timeStart)
		statusCode := wrappedRW.statusCode

		body := string(wrappedRW.body)

		var logEvent *zerolog.Event
		if statusCode >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		logEvent.
			Str("method", r.Method).
			Int("status_code", statusCode).
			Str("path", path).
			Str("latency", latency.String()).
			Str("request_id", requestId).
			Msg(body)
	})
}
