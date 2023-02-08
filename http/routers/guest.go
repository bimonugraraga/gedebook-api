package routers

import (
	"gedebook.com/api/controllers"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
)

func Guest(g *gin.RouterGroup) {
	rp := repository.InitRepositoryInstance()
	userSrv := services.NewUserService(rp.User)
	bookSrv := services.NewBookService(rp.Book, rp.User, rp.Category)

	guestCtl := controllers.NewGuestController(userSrv, bookSrv)

	//!Book
	g.GET("/book/:id", guestCtl.GetOneBook)
}
