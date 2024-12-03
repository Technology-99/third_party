package jwts

import "errors"

var (
	ErrorTokenInvalidAudience = errors.New("invalid audience")
	ErrorTokenHasExpired      = errors.New("token has expired")

	ErrorTokenInvalidIssuer = errors.New("invalid issuer")

	ErrorTokenNotActiveYet = errors.New("token not active yet")
)
