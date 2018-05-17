package main

import (
	"mugg/gcat/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.API(r)

	r.Run(":8989")
}
