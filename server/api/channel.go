package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/vitalis-virtus/video-chat/models"
)

var upg = websocket.Upgrader{
	ReadBufferSize:  102400,
	WriteBufferSize: 102400,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// CreateChannel godoc
// @Summary Create new channel
// @Description Create new channel.
// @Tags channels
// @Produce json
// @Success 200 {object} models.CreateChannelRes "Success"
// @Failure 500 {object} models.HTTPError "Cannot create channel"
// @Router /channels [post]
func (api *api) CreateChannel(c *gin.Context) {
	id := api.services.CreateChannel()

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// JoinChannel godoc
// @Summary Join to channet
// @Description Join to channel.
// @Tags channels
// @Param id path int true "Channel ID"
// @Success 101 {string} string "WebSocket Protocol Switch"
// @Failure 500 {object} models.HTTPError "Internal server error"
// @Router /channels/{id} [get]
func (api *api) JoinChannel(c *gin.Context) {
	params := new(models.JoinChannelQuery)
	// if err := c.ShouldBindQuery(&params); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "missed participant name",
	// 	})
	// 	return
	// }

	participantIP := c.ClientIP()

	if participantIP == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missed participant IP",
		})
		return
	}

	params.IP = participantIP

	var req models.UriIDString
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missed channel ID",
		})
		return
	}

	params.ChannelID = req.ID

	conn, err := upg.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	api.services.JoinChannel(conn, params)
}
