package errs

import (
	"net/http"

	"gedebook.com/api/dto/responses"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context, code int, customMsg string) {
	switch code {
	case 400:
		ctx.JSON(http.StatusBadRequest, responses.R{
			Code:    http.StatusBadRequest,
			Message: customMsg,
		})
	case 401:
		ctx.JSON(http.StatusUnauthorized, responses.R{
			Code:    http.StatusUnauthorized,
			Message: customMsg,
		})
	case 403:
		ctx.JSON(http.StatusForbidden, responses.R{
			Code:    http.StatusForbidden,
			Message: customMsg,
		})
	case 404:
		ctx.JSON(http.StatusNotFound, responses.R{
			Code:    http.StatusNotFound,
			Message: customMsg,
		})
	default:
		ctx.JSON(http.StatusInternalServerError, responses.R{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}
}
