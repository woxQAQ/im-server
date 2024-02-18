package jwt

import "github.com/golang-jwt/jwt/v4"

func GetToken(secretkey string, iat, second, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + second
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretkey))
}
