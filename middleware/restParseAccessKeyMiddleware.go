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
		// TODO generate middleware implement function, delete after code implementation
		accessKey := r.Header.Get(commKey.HANDER_ACCESSKEY)
		ctx := context.WithValue(r.Context(), "AccessKey", accessKey)
		r = r.WithContext(ctx)
		// Passthrough to next handler if need
		next(w, r)
	}
}
