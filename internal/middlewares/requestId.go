package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	requestID = "X-Request_ID"
)

//receive handler and return handler
func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(requestID)
		if id == "" {
			id = uuid.NewString()
		}

		w.Header().Add("requestId", id)
		ctx := context.WithValue(r.Context(), "requestCtxId", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}