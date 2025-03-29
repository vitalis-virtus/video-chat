package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitalis-virtus/video-chat/config"
)

func (api *api) Index(c *gin.Context) {
	c.String(http.StatusOK, "This is a service '%s'", config.Service)
}

func (api *api) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
