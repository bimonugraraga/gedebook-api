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
	// g.GET("/profile", middlewares.UserAuthn(), func(c *gin.Context) {
	// 	value, exists := c.Get("user")
	// 	myClaim, ok := value.(constants.AuthnPayload)
	// 	fmt.Println(myClaim.ID, ok)
	// 	fmt.Println(value, exists)
	// 	c.JSON(http.StatusOK, responses.R{
	// 		Code:    http.StatusOK,
	// 		Message: "Hello World",
	// 	})
	// })
	g.GET("/profile", middlewares.UserAuthn(), userCtl.UserProfile)
}
