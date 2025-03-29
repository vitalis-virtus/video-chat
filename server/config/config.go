package config

import (
	"os"
	"strconv"
	"strings"
)

type (
	Config struct{ API }
	API    struct {
		ListenPort         uint64
		CORSAllowedOrigins []string
	}
)

const (
	Service = "video-chat"
)

func GetNewConfig() (*Config, error) {
	port, err := strconv.ParseUint(os.Getenv("LISTEN_PORT"), 10, 64)
	if err != nil {
		return nil, err
	}

	cfg := Config{
		API: API{
			ListenPort:         port,
			CORSAllowedOrigins: strings.Split(os.Getenv("CORS_ALLOWED"), ","),
		},
	}

	return &cfg, nil
}
