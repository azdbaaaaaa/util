package xjwt

import (
	"github.com/dgrijalva/jwt-go"
)

func GenToken(algorithm jwt.SigningMethod, secret string, claims jwt.MapClaims) (token string, err error) {
	at := jwt.NewWithClaims(algorithm, claims)
	token, err = at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(string), nil
}
