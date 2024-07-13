package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type PathHttpInterceptorMiddleware struct {
}

func NewPathHttpInterceptorMiddleware() *PathHttpInterceptorMiddleware {
	return &PathHttpInterceptorMiddleware{}
}

func (m *PathHttpInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "FullMethod", r.URL.Path)
		ctx = context.WithValue(ctx, "RequestURI", r.RequestURI)
		fullAddr := httpx.GetRemoteAddr(r)
		fullAddrAndPort := strings.Split(fullAddr, ":")
		logx.Infof("client ip : %s", fullAddrAndPort[0])
		logx.Infof("client port : %s", fullAddrAndPort[1])
		ctx = context.WithValue(r.Context(), "clientIp", fullAddrAndPort[0])
		ctx = context.WithValue(r.Context(), "clientPort", fullAddrAndPort[1])
		r = r.WithContext(ctx)
		next(w, r)
	}
}
