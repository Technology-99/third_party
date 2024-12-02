package jwts

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtRsaCommonCustomClaims struct {
	LoginTime time.Time
	jwt.StandardClaims
}

func JwtRsaCommonCreateToken(exp int, privateKey string, loginTime time.Time) (string, int64, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", -1, err
	}
	expiresAt := time.Now().Add(time.Duration(exp) * time.Second).Unix()
	customClaims := &JwtRsaCommonCustomClaims{
		LoginTime: loginTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt, // 过期时间
		},
	}
	//采用 RS256 加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, customClaims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", -1, err
	}
	return tokenString, expiresAt, nil
}

// 解析 token
func JwtRsaCommonParseToken(tokenString, pubKey string) (*JwtRsaCommonCustomClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenString, &JwtRsaCommonCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if claims, ok := token.Claims.(*JwtRsaCommonCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
