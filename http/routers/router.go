package routers

import (
	"net/http"

	"gedebook.com/api/controllers"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/dto/responses"
	"gedebook.com/api/services"
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
	rp := repository.InitRepositoryInstance()
	adminSrv := services.NewAdminService(rp.Admin)
	adminCtl := controllers.NewAdminController(adminSrv)

	admin := r.Group("/admin")
	{
		admin.GET("/test",
			adminCtl.HelloAdmin,
		)
	}
}
