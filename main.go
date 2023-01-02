package main

import (
	"log"

	"gedebook.com/api/db"
	"gedebook.com/api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.RedirectFixedPath = true
	routers.RoutesHandler(r)
	db.InitDB()
	err := r.Run(":3000")
	if err != nil {
		log.Fatal("Error To Start")
	}
}
