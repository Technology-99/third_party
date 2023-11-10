/*
*

	@author: taco
	@Date: 2023/9/27
	@Time: 10:16

*
*/
package middleware

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"third_party/cache_key"
	"third_party/commKey"
	utils "third_party/cryptography"
)

type JwtVerifyMiddleware struct {
	SvcName string
	Redis   *redis.Redis
	Rbac    *casbin.Enforcer
}

func NewJwtVerifyMiddleware(name string, rdb *redis.Redis, rbac *casbin.Enforcer) *JwtVerifyMiddleware {
	return &JwtVerifyMiddleware{
		SvcName: name,
		Redis:   rdb,
		Rbac:    rbac,
	}
}

func (m *JwtVerifyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		authToken := r.Header.Get(commKey.HANDER_AUTHORIZATION)
		if authToken == "" || len(authToken) <= 7 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token is empty"))
			return
		}
		token := authToken[7:]
		key := fmt.Sprintf(cache_key.ACCESS_TOKEN_KEY, m.SvcName, r.Header.Get(commKey.HANDER_ACCESSKEY))
		pubKey, err := m.Redis.Get(key)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("accessKey is empty"))
			return
		}
		claims, err := utils.ParseToken(token, pubKey)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token is invalid"))
			return
		}
		//权限验证
		if claims.UserId != 1 {
			ok, err := m.checkPermission(fmt.Sprintf("%d", claims.UserId), fmt.Sprintf("%d", claims.TenantID), m.SvcName, r.RequestURI)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("checkPermission is invalid"))
				return
			}
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Permission verification failed"))
				return
			}
		}

		ctx := context.WithValue(r.Context(), commKey.CONTEXT_KEY_UID, claims.UserId)
		ctx = context.WithValue(ctx, commKey.CONTEXT_KEY_TENANTID, claims.TenantID)
		r = r.WithContext(ctx)
		// Passthrough to next handler if need
		next(w, r)
	}
}

func (m *JwtVerifyMiddleware) checkPermission(sub, domain, obj, act string) (bool, error) {
	ok, err := m.Rbac.Enforce(sub, domain, obj, act)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	return true, nil
}
