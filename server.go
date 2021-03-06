package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kzdv/overflight-api/middleware"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(appenv string) *Server {
	server := Server{}

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger)
	server.engine = engine
	engine.LoadHTMLGlob("static/*")

	SetupRoutes(engine)

	return &server
}
