package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Jwt interface {
	GetToken(secretkey string, info any) (string, error)
}

type Claims struct {
	info any
	jwt.RegisteredClaims
}

func (c *Claims) GetToken(secretkey string, info any) (string, error) {
	claims := Claims{
		info: info,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(time.Second.Microseconds() * accessExpired))),
			Issuer:    "woxQAQ",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretkey))
}

type ClaimsWithUserId struct {
	UserId string
	jwt.RegisteredClaims
}

type ClaimsWithEmail struct {
	Email string
	jwt.RegisteredClaims
}

type ClaimsWithPhone struct {
	Phone string
	jwt.RegisteredClaims
}

func GetTokenWithPhone(secretkey string, phone string, accessExpired int64) (string, error) {
	claims := ClaimsWithPhone{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(time.Second.Microseconds() * accessExpired))),
			Issuer:    "woxQAQ",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretkey))
}

func GetTokenWithEmail(secretkey string, email string) (string, error) {
	claims := ClaimsWithEmail{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "woxQAQ",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretkey))
}

func GetTokenWithUid(secretkey string, UserId string) (string, error) {
	claims := ClaimsWithUserId{
		UserId: UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "woxQAQ",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretkey))
}
