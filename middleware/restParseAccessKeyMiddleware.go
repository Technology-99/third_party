package middleware

import (
	"context"
	"github.com/Technology-99/third_party/commKey"
	"net/http"
)

type ParseAccessKeyMiddleware struct {
}

func NewParseAccessKeyMiddleware() *ParseAccessKeyMiddleware {
	return &ParseAccessKeyMiddleware{}
}

func (m *ParseAccessKeyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessKey := r.Header.Get(commKey.HeaderXAccessKeyFor)
		ctx := context.WithValue(r.Context(), CtxXAccessKeyFor, accessKey)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
