package utils

import (
	"gedebook.com/api/domain"
	"gedebook.com/api/env"
	"github.com/kataras/jwt"
)

type UserClaim struct {
	Username string `json:"username"`
}

func SignToken(payload domain.UserPayload) (string, error) {
	c := env.Init()
	sharedKey := []byte(c.SecretKey)
	token, err := jwt.Sign(jwt.HS256, sharedKey, payload)
	return string(token), err
}
