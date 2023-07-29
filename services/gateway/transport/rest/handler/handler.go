package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rvinnie/bigO/services/gateway/config"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.GIN.Mode)

	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api := router.Group("/api")
	{
		h.initComplexityRoutes(api)
		h.initDonatesRoutes(api)
	}

	return router
}
