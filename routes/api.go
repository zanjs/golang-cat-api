package routes

import (
	"mugg/gcat/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// API is router api
func API(g *gin.Engine) {
	catController := controllers.NewCat()

	v1 := g.Group("/v1")
	{
		v1.GET("/", catController.Home)

		v1.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
	}

}
