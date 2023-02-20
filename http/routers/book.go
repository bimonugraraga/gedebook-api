package routers

import (
	"gedebook.com/api/controllers"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/middlewares"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
)

func Book(g *gin.RouterGroup) {
	rp := repository.InitRepositoryInstance()
	bookSrv := services.NewBookService(rp.Book, rp.User, rp.Category)

	bookCtl := controllers.NewBookController(bookSrv)

	g.GET("", middlewares.UserAuthn(), bookCtl.GetAllBook)
	g.GET("/:book_id", middlewares.UserAuthn(), bookCtl.GetOneBook)
	g.POST("", middlewares.UserAuthn(), bookCtl.CreateBook)
	g.PUT("/:book_id", middlewares.UserAuthn(), bookCtl.UpdateBook)
}
