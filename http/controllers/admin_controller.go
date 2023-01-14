package controllers

import (
	"net/http"

	"gedebook.com/api/dto/responses"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
)

type AdminController interface {
	//!Attach Your Function Here
	HelloAdmin(ctx *gin.Context)
}

type adminController struct {
	adminSrv services.AdminService
}

func NewAdminController(adminSrv services.AdminService) AdminController {
	return &adminController{
		adminSrv: adminSrv,
	}
}

//!Code Your Function Here
func (ctl *adminController) HelloAdmin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responses.R{
		Code:    http.StatusOK,
		Message: "Hello Admin",
	})
}
