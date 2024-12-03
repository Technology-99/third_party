/*
*

	@author: howard
	@Date: 2023/12/03
	@Time: 10:16

*
*/
package middleware

import (
	"context"
	"fmt"
	"github.com/Technology-99/third_party/cache_key"
	"github.com/Technology-99/third_party/commKey"
	"github.com/Technology-99/third_party/jwts"
	"github.com/Technology-99/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type JwtRSACommonVerifyMiddleware struct {
	AuthServiceName string
	PublicKey       string
	Redis           *redis.Redis
}

func NewJwtRSAVerifyMiddleware(authServiceName string, PublicKey string, rdb *redis.Redis) *JwtRSACommonVerifyMiddleware {
	return &JwtRSACommonVerifyMiddleware{
		AuthServiceName: authServiceName,
		PublicKey:       PublicKey,
		Redis:           rdb,
	}
}
func (m *JwtRSACommonVerifyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get(commKey.HANDER_AUTHORIZATION)
		if authToken == "" || len(authToken) <= 7 {
			CommonErrResponse(w, r, response.AUTHORIZATION_NOT_FOUND)
			return
		}
		token := authToken[7:]
		key := fmt.Sprintf(cache_key.ACCESS_TOKEN_KEY, m.AuthServiceName, r.Header.Get(commKey.HANDER_ACCESSKEY))
		logx.Infof("key: %v", key)
		cacheAccessToken, err := m.Redis.Get(key)
		if err != nil {
			logx.Errorf("cacheAccessToken error: %v", err)
			CommonErrResponse(w, r, response.ACCESSKEY_NOT_FOUND)
			return
		}
		if cacheAccessToken == "" || len(cacheAccessToken) <= 0 || cacheAccessToken != token {
			logx.Infof("this accessToken is expired")
			CommonErrResponse(w, r, response.ACCESS_EXPIRED)
			return
		}
		claims, err := jwts.JwtRSACommonParseAndVerifyToken(token, m.PublicKey)
		if err != nil {
			logx.Errorf("ParseToken error: %v", err)
			CommonErrResponse(w, r, response.ACCESS_EXPIRED)
			return
		}
		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
