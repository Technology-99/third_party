package middleware

import (
	"context"
	"github.com/Technology-99/third_party/sony"
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
		ctx := context.WithValue(r.Context(), CtxFullMethod, r.URL.Path)
		ctx = context.WithValue(ctx, CtxRequestURI, r.RequestURI)
		fullAddr := httpx.GetRemoteAddr(r)
		fullAddrAndPort := strings.Split(fullAddr, ":")
		ctx = context.WithValue(ctx, CtxClientIp, fullAddrAndPort[0])
		logx.Infof("client ip : %s", fullAddrAndPort[0])

		requestID := ""
		if ctx.Value(CtxRequestID) == nil {
			requestID = sony.NextId()
		} else {
			requestID = ctx.Value(CtxRequestID).(string)
		}
		ctx = context.WithValue(ctx, CtxRequestID, requestID)
		// 获取 User-Agent
		userAgent := r.Header.Get(CtxUserAgent)
		ctx = context.WithValue(ctx, CtxUserAgent, userAgent)
		//ctx = context.WithValue(ctx, "clientPort", fullAddrAndPort[1])
		r = r.WithContext(ctx)
		next(w, r)
	}
}
