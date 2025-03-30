package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/vitalis-virtus/video-chat/docs"

	"github.com/joho/godotenv"
	"github.com/vitalis-virtus/video-chat/api"
	"github.com/vitalis-virtus/video-chat/config"
	"github.com/vitalis-virtus/video-chat/services"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

//	@title			vide-chat API
//	@version		1.0
//	@description	This is a swagger specification for video-chat server backend.

// @host						localhost
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				    Used for secure private routes
//
// @BasePath
func main() {
	cfg, err := config.GetNewConfig()
	if err != nil {
		log.Fatal("can`t read config from file ", err)
	}

	s := services.New(&cfg.Services)

	api := api.New(&cfg.API, s)
	if err != nil {
		log.Fatal("api.New ", err)
	}

	go func() {
		err := api.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulStop

	result := make(chan error)
	go func() {
		result <- api.Stop()
	}()

	select {
	case err := <-result:
		log.Fatal(err)
	case <-time.After(time.Second * 15):
		log.Println("stoped by timeout")
	}
}
