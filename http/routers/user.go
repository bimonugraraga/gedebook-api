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

	bookSrv := services.NewBookService(rp.Book, rp.User, rp.Category)
	bookCtl := controllers.NewBookController(bookSrv)

	//!Auth
	g.POST("/register", userCtl.UserRegister)
	g.POST("/login", userCtl.UserLogin)

	//!Profile
	g.GET("/profile", middlewares.UserAuthn(), userCtl.UserProfile)

	//!Book
	g.POST("/book", middlewares.UserAuthn(), bookCtl.CreateBook)
	g.PUT("/book/:id", middlewares.UserAuthn(), bookCtl.UpdateBook)
}
