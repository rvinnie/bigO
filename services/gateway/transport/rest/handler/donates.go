package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initDonatesRoutes(api *gin.RouterGroup) {
	donates := api.Group("/donates")
	{
		donates.GET("stub", h.donatesStub)
	}
}

func (h *Handler) donatesStub(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"result": "donates stub",
	})
}
