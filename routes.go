package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzdv/overflight-api/api"
)

func SetupRoutes(engine *gin.Engine) {
	engine.StaticFile("/", "./static/index.html")
	engine.StaticFile("/openapi.yaml", "./static/openapi.yaml")

	engine.GET("/live/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PONG"})
	})

	engine.GET("/live/:fac", api.GetLive)
}
