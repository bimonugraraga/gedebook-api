package routers

import (
	"net/http"

	"gedebook.com/api/dto/responses"
	"github.com/gin-gonic/gin"
)

type GinEng gin.Engine

func RoutesHandler(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Hello World",
		})
	})
}
