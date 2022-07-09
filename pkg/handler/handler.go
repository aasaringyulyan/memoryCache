package handler

import (
	"github.com/gin-gonic/gin"
	"memoryCache/pkg/service"
	"time"
)

type Config struct {
	Duration time.Duration
}

type Handler struct {
	services *service.Service
	config   Config
}

func NewHandler(services *service.Service, cfg Config) *Handler {
	return &Handler{
		services: services,
		config:   cfg,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/add", h.setItem)
	router.GET("/get", h.getValue)
	router.DELETE("/delete", h.deleteItem)
	router.GET("/search", h.searchKey)
	router.GET("/increase", h.increaseValue)
	router.GET("/reduce", h.reduceValue)

	return router
}
