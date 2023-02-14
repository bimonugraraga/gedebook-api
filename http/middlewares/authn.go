package middlewares

import (
	"strings"

	"gedebook.com/api/constants"
	"gedebook.com/api/errs"
	"gedebook.com/api/utils"
	"github.com/gin-gonic/gin"
)

func UserAuthn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header.Values("Authorization")) == 0 {
			ctx.Next()
			return
		}
		access_token := strings.Split(ctx.Request.Header.Values("Authorization")[0], " ")
		if access_token[0] != "Bearer" {
			errs.ErrorHandler(ctx, 401, "Invalid Access Token")
			return
		}

		payload, err := utils.VerifyToken(access_token[1])
		if err != nil {
			errs.ErrorHandler(ctx, 401, "Invalid Access Token")
			return
		}
		setPayload := constants.AuthnPayload{
			ID:    payload.ID,
			Email: payload.Email,
			Name:  payload.Name,
			Role:  payload.Role,
		}
		switch payload.Role {
		case string(constants.User):
			ctx.Set("user", setPayload)
		case string(constants.Admin):
			ctx.Set("admin", setPayload)
		default:
			errs.ErrorHandler(ctx, 401, "Invalid Access Token")
			return
		}
		ctx.Next()
	}
}
