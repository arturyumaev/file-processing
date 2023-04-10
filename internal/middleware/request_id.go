package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	ContextKeyRequestID = "requestId"
	RequestIdHeaderName = "X-Request-ID"
)

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.New().String()
		nextCtx := context.WithValue(r.Context(), ContextKeyRequestID, requestId)
		r = r.WithContext(nextCtx)
		w.Header().Set(RequestIdHeaderName, requestId)

		next.ServeHTTP(w, r)
	})
}
