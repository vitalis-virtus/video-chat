package services

import "github.com/vitalis-virtus/video-chat/config"

type Services interface {
	CreateChannel() string
}

type service struct {
	cfg   *config.Services
	rooms Rooms
}

func New(cfg *config.Services) Services {
	rooms := NewRoom()

	rooms.Init()

	return &service{
		cfg:   cfg,
		rooms: rooms,
	}
}
