package routers

import (
	"gedebook.com/api/controllers"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/middlewares"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
)

func User(g *gin.RouterGroup) {
	rp := repository.InitRepositoryInstance()
	userSrv := services.NewUserService(rp.User)
	userCtl := controllers.NewUserController(userSrv)

	//!Auth
	g.POST("/register", userCtl.UserRegister)
	g.POST("/login", userCtl.UserLogin)

	//!Profile
	g.GET("/profile", middlewares.UserAuthn(), userCtl.UserProfile)
	g.PUT("/profile", middlewares.UserAuthn(), userCtl.UpdateProfile) //!Cannot Change Email And Password

}
