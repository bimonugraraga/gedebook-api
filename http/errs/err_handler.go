package errs

import (
	"net/http"
	"regexp"
	"strings"

	"gedebook.com/api/dto/responses"
	"github.com/gin-gonic/gin"
)

func ParseSQLError(err error) (errNum string) {
	parsed := strings.Split(err.Error(), " ")
	regex, _ := regexp.Compile(`SQLSTATE`)
	var isMatch = regex.MatchString(parsed[len(parsed)-1])
	if isMatch {
		regex, _ = regexp.Compile(`[0-9]+`)
		errNum = regex.FindString(parsed[len(parsed)-1])
		return errNum
	}
	return ""
}
func ErrorHandler(ctx *gin.Context, code int, customMsg string) {
	switch code {
	case 400:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.R{
			Code:    http.StatusBadRequest,
			Message: customMsg,
		})
	case 401:
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.R{
			Code:    http.StatusUnauthorized,
			Message: customMsg,
		})
	case 403:
		ctx.AbortWithStatusJSON(http.StatusForbidden, responses.R{
			Code:    http.StatusForbidden,
			Message: customMsg,
		})
	case 404:
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.R{
			Code:    http.StatusNotFound,
			Message: customMsg,
		})
	default:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.R{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}
}
