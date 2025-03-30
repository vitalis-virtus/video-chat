package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitalis-virtus/video-chat/config"
)

// Index godoc
// @Summary Return a message about the service
// @Description Return a message about the service using the service name from configuration
// @Tags common
// @Produce plain
// @Success 200 {string} string "This is a service 'service_name'"
// @Router / [get]
func (api *api) Index(c *gin.Context) {
	c.String(http.StatusOK, "This is a service '%s'", config.Service)
}

// Health godoc
// @Summary Return a health check response
// @Description Return a success response indicating that the service is healthy
// @Tags common
// @Produce json
// @Success 200 {object} models.HTTPSuccess "All good"
// @Router /health [get]
func (api *api) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
