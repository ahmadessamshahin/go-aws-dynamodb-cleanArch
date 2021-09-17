package jwtparser

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	IssuedAt   int64  `json:"iat"`
	Expiration int64  `json:"exp"`
	Username   string `json:"Username"`
}

func Validate(tokenString, secret string) (claim Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || token.Valid {
		return claim, fmt.Errorf("unable to parse the token : message: %s", err)
	}

	return claim, nil
}

func (c Claims) Valid() error {

	if c.Expiration == 0 || c.Username == "" {
		return errors.New("missing jet fields")
	}
	now := time.Now()
	exp := time.Unix(c.Expiration, 0)
	if now.After(exp) {
		return errors.New("token is expired")
	}
	return nil
}
