package utils

import (
	"gedebook.com/api/constants"
	"gedebook.com/api/env"
	"github.com/kataras/jwt"
)

type UserClaim struct {
	Username string `json:"username"`
}

func SignToken(payload constants.AuthnPayload) (string, error) {
	c := env.Init()
	sharedKey := []byte(c.SecretKey)
	token, err := jwt.Sign(jwt.HS256, sharedKey, payload)
	return string(token), err
}

func VerifyToken(access_token string) (*constants.AuthnPayload, error) {
	c := env.Init()
	sharedKey := []byte(c.SecretKey)
	token := []byte(access_token)
	verifiedToken, err := jwt.Verify(jwt.HS256, sharedKey, token)
	resp := new(constants.AuthnPayload)
	err = verifiedToken.Claims(&resp)
	return resp, err
}
