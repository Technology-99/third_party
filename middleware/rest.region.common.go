package middleware

import (
	"context"
	"github.com/Technology-99/third_party/sony"
	"github.com/lionsoul2014/ip2region/v1.0/binding/golang/ip2region"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
	"time"
)

type RegionInterceptorMiddleware struct {
	Region *ip2region.Ip2Region
}

func NewRegionInterceptorMiddleware(region *ip2region.Ip2Region) *RegionInterceptorMiddleware {
	return &RegionInterceptorMiddleware{
		Region: region,
	}
}

func (m *RegionInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		clientIp := ""
		if ctx.Value("ClientIp") == nil {
			clientIp = sony.NextId()
		} else {
			fullAddr := httpx.GetRemoteAddr(r)
			fullAddrAndPort := strings.Split(fullAddr, ":")
			ctx = context.WithValue(ctx, "ClientIp", fullAddrAndPort[0])
			logx.Infof("client ip : %s", fullAddrAndPort[0])
			clientIp = fullAddrAndPort[0]
		}

		startTime := time.Now()
		logc.Debugf(ctx, "起始时间: %s", startTime.Format(time.DateTime))

		info, _ := m.Region.MemorySearch(clientIp)
		ctx = context.WithValue(ctx, "cityId", info.CityId)
		ctx = context.WithValue(ctx, "country", info.Country)
		ctx = context.WithValue(ctx, "region", info.Region)
		ctx = context.WithValue(ctx, "province", info.Province)
		ctx = context.WithValue(ctx, "city", info.City)
		ctx = context.WithValue(ctx, "iSP", info.ISP)
		endTime := time.Now()
		logc.Debugf(ctx, "结束时间: %s", endTime.Format(time.DateTime))
		logc.Infof(ctx, "地理位置中间件耗时: %v", endTime.Sub(startTime).Milliseconds())

		r = r.WithContext(ctx)
		next(w, r)
	}
}
