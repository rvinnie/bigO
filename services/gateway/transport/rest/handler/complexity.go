package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initComplexityRoutes(api *gin.RouterGroup) {
	complexity := api.Group("/complexity")
	{
		complexity.GET("/stub", h.complexityStub)
	}
}

func (h *Handler) complexityStub(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"result": "complexity stub",
	})
}
