package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitalis-virtus/video-chat/server/config"
	"github.com/vitalis-virtus/video-chat/server/services"
)

type API interface {
}

type api struct {
	cfg  config.Config
	router *gin.Engine
	server *http.Server
	services  services.Services
}