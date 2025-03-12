package qxJwts

import (
	"github.com/golang-jwt/jwt/v4"
)

func JwtParseUnverified(token string) (*jwt.MapClaims, error) {
	parseToken, _, err := jwt.NewParser().ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return nil, ErrorTokenInvalid
	}

	// 获取 Payload（Claims）
	tempclaims, parseJwtClaimsOk := parseToken.Claims.(jwt.MapClaims)
	if !parseJwtClaimsOk {
		return nil, ErrorJwtClaimsInvalid
	}
	return &tempclaims, nil
}
