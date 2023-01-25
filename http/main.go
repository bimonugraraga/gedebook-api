package main

import (
	"log"

	"gedebook.com/api/db"
	"gedebook.com/api/env"
	"gedebook.com/api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	c := env.Init()
	dbInstance := db.InitDB(&c.Database)
	defer dbInstance.Close()
	r := gin.New()
	r.RedirectFixedPath = true
	routers.RoutesHandler(r)
	err := r.Run(":3000")
	if err != nil {
		log.Fatal("Error To Start")
	}
}
