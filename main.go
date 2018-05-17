package main

import (
	"mugg/gcat/app/conf"
	"mugg/gcat/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := conf.Config
	routes.API(r)

	r.Run(":" + config.APP.Port)
}
