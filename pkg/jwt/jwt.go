package jwt

import (
	"fiber-wire-template/pkg/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
	key []byte
}

type Claims struct {
	UserId string
	jwt.RegisteredClaims
}

func NewJwt(conf *config.Config) *JWT {
	return &JWT{key: []byte(conf.AppCnf.JWTSecret)}
}
func (j *JWT) GenerateToken(userId string, expireTime time.Time) (string, error) {
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			ID:        "",
			Audience:  []string{},
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return token, err
}

func (j *JWT) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
