package services

import (
	"github.com/gorilla/websocket"
	"github.com/vitalis-virtus/video-chat/config"
	"github.com/vitalis-virtus/video-chat/models"
)

type Services interface {
	CreateChannel() string
	JoinChannel(conn *websocket.Conn, joinParams *models.JoinChannelQuery)
}

type service struct {
	cfg       *config.Services
	rooms     Rooms
	broadcast chan broadcastMsg
}

func New(cfg *config.Services) Services {
	rooms := NewRoom()

	rooms.Init()

	return &service{
		cfg:       cfg,
		rooms:     rooms,
		broadcast: make(chan broadcastMsg),
	}
}
