package routers

import (
	"gedebook.com/api/controllers"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/middlewares"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
)

func Chapter(g *gin.RouterGroup) {
	rp := repository.InitRepositoryInstance()
	chapterSrv := services.NewChapterService(rp.Chapter, rp.Book, rp.User)
	chapterCtl := controllers.NewChapterController(chapterSrv)

	g.POST("", middlewares.UserAuthn(), chapterCtl.CreateChapter)
	g.PUT("/:chapter_id", middlewares.UserAuthn(), chapterCtl.UpdateChapter)
	g.GET("/:chapter_id/book/:book_id", middlewares.UserAuthn(), chapterCtl.TestChapter)
}
