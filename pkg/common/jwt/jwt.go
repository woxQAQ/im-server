package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	info any
	jwt.RegisteredClaims
}

var ErrTokenInvalid = errors.New("TokenInvalid")

type Jwt struct {
	SignedKeys   []byte
	AccessExpire int64
}

func (j *Jwt) GetToken(info any) (string, error) {
	claims := Claims{
		info: info,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Duration(time.Second.Microseconds() * j.AccessExpire)),
			),
			Issuer:   "woxQAQ",
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignedKeys)
}

func (j *Jwt) GetTokenWithClaims(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignedKeys)
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Duration(time.Second.Microseconds() * accessExpired)),
			),
			Issuer:   "woxQAQ",
			IssuedAt: jwt.NewNumericDate(time.Now()),
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

func (j *Jwt) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignedKeys, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, jwt.ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, jwt.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, jwt.ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

func (j *Jwt) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignedKeys, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		return j.GetTokenWithClaims(claims)
	}

}
