package middleware

import (
	"context"
	"github.com/ua-parser/uap-go/uaparser"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
	"time"
)

type UaParserInterceptorMiddleware struct {
	Uaparser *uaparser.Parser
}

func NewUaParserInterceptorMiddleware(uaparser *uaparser.Parser) *UaParserInterceptorMiddleware {
	return &UaParserInterceptorMiddleware{
		Uaparser: uaparser,
	}
}

func (m *UaParserInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		userAgent := ""
		if ctx.Value("User-Agent") == nil {
			fullAddr := httpx.GetRemoteAddr(r)
			fullAddrAndPort := strings.Split(fullAddr, ":")
			ctx = context.WithValue(ctx, "clientIp", fullAddrAndPort[0])
			logx.Infof("client ip : %s", fullAddrAndPort[0])
			userAgent = fullAddrAndPort[0]
		} else {
			userAgent = ctx.Value("UserAgent").(string)
		}

		startTime := time.Now()
		logc.Debugf(ctx, "起始时间: %s", startTime.Format(time.DateTime))

		client := m.Uaparser.Parse(userAgent)
		ctx = context.WithValue(ctx, "UserAgent.Family", client.UserAgent.Family)
		ctx = context.WithValue(ctx, "UserAgent.Major", client.UserAgent.Major)
		ctx = context.WithValue(ctx, "UserAgent.Minor", client.UserAgent.Minor)
		ctx = context.WithValue(ctx, "UserAgent.Patch", client.UserAgent.Patch)

		ctx = context.WithValue(ctx, "Os.Family", client.Os.Family)
		ctx = context.WithValue(ctx, "Os.Major", client.Os.Major)
		ctx = context.WithValue(ctx, "Os.Minor", client.Os.Minor)
		ctx = context.WithValue(ctx, "Os.Patch", client.Os.Patch)
		ctx = context.WithValue(ctx, "Os.PatchMinor", client.Os.PatchMinor)

		ctx = context.WithValue(ctx, "Device.Family", client.Device.Family)
		ctx = context.WithValue(ctx, "Device.Brand", client.Device.Brand)
		ctx = context.WithValue(ctx, "Device.Model", client.Device.Model)
		endTime := time.Now()
		logc.Debugf(ctx, "结束时间: %s", endTime.Format(time.DateTime))
		logc.Infof(ctx, "设备解析中间件耗时: %v", endTime.Sub(startTime).Milliseconds())

		r = r.WithContext(ctx)
		next(w, r)
	}
}
