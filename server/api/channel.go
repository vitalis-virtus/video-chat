package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *api) CreateChannel(c *gin.Context) {
	id := api.services.CreateChannel()

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
