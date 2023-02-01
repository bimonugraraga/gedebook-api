package routers

import (
	"gedebook.com/api/controllers"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
)

func Admin(g *gin.RouterGroup) {
	rp := repository.InitRepositoryInstance()
	adminSrv := services.NewAdminService(rp.Admin)
	adminCtl := controllers.NewAdminController(adminSrv)

	g.GET("/check", adminCtl.HelloAdmin)
}
