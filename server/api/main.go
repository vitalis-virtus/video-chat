package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vitalis-virtus/video-chat/config"
	"github.com/vitalis-virtus/video-chat/services"
	"go.uber.org/zap"
)

type API interface {
	Run() error
	Stop() error
}

type api struct {
	cfg      *config.API
	router   *gin.Engine
	server   *http.Server
	services services.Services
}

func New(cfg *config.API, s services.Services) API {
	api := api{
		cfg:      cfg,
		services: s,
	}

	api.initialize()

	return &api
}

func (api *api) Run() error {
	return api.startServe()
}

func (api *api) Stop() error {
	return api.server.Shutdown(context.Background())
}

func (api *api) initialize() {
	api.router = gin.Default()

	api.router.Use(gin.Logger())
	api.router.Use(gin.Recovery())

	api.router.Use(cors.New(cors.Config{
		AllowOrigins:     api.cfg.CORSAllowedOrigins,
		AllowCredentials: true,
		AllowMethods: []string{
			http.MethodPost, http.MethodHead, http.MethodGet, http.MethodOptions, http.MethodPut, http.MethodDelete,
		},
		AllowHeaders: []string{
			"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token",
			"Authorization", "User-Env", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Access-Control-Max-Age",
		},
		MaxAge: time.Second * 86400,
	}))

	api.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api.router.GET("/", api.Index)
	api.router.GET("/health", api.Health)

	channelsGroup := api.router.Group("/channels")

	{
		channelsGroup.POST("", api.CreateChannel)
		channelsGroup.GET("/:id", api.JoinChannel)
	}

	api.server = &http.Server{Addr: fmt.Sprintf(":%d", api.cfg.ListenPort), Handler: api.router}
}

func (api *api) startServe() error {
	log.Println("Start listening server on port", zap.Uint64("port", api.cfg.ListenPort))

	err := api.server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("API server was closed")
			return nil
		}

		return fmt.Errorf("cannot run API service: %s", err.Error())
	}

	return nil
}
