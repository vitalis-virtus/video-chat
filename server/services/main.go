package services

import "github.com/vitalis-virtus/video-chat/config"

type Services interface{}

type service struct {
	cfg *config.Services
}

func New(cfg *config.Services) Services {
	return &service{}
}
