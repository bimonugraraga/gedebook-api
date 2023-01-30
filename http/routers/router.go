package routers

import (
	"fmt"
	"net/http"

	"gedebook.com/api/constants"
	"gedebook.com/api/controllers"
	"gedebook.com/api/db"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/dto/responses"
	"gedebook.com/api/middlewares"
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
	if db.GetConn() == nil {
		panic("SINI")
	}
	adminSrv := services.NewAdminService(rp.Admin)
	adminCtl := controllers.NewAdminController(adminSrv)

	userSrv := services.NewUserService(rp.User)
	userCtl := controllers.NewUserController(userSrv)

	admin := r.Group("/admin")
	{
		admin.GET("/test",
			adminCtl.HelloAdmin,
		)
	}

	user := r.Group("/user")
	{
		user.POST("/register",
			userCtl.UserRegister,
		)
		user.POST("/login",
			userCtl.UserLogin,
		)
		user.GET("/profile", middlewares.UserAuthn(), func(c *gin.Context) {
			value, exists := c.Get("user")
			myClaim, ok := value.(constants.AuthnPayload)
			fmt.Println(myClaim.ID, ok)
			fmt.Println(value, exists)
			c.JSON(http.StatusOK, responses.R{
				Code:    http.StatusOK,
				Message: "Hello World",
			})
		})
	}
}
