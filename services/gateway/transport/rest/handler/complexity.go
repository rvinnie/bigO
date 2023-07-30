package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/rvinnie/bigO/services/gateway/pb"
	"net/http"
)

const (
	MaxRequestBuffer = 4096
)

func (h *Handler) initComplexityRoutes(api *gin.RouterGroup) {
	complexity := api.Group("/complexity")
	{
		complexity.GET("/stub", h.complexityStub)
		complexity.POST("/count", h.countAlgorithmComplexity)
	}
}

func (h *Handler) complexityStub(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"result": "complexity stub",
	})
}

type AlgorithmRequestBody struct {
	Algorithm string `json:"algorithm" binding:"required"`
}

func (h *Handler) countAlgorithmComplexity(c *gin.Context) {
	var algorithmRequestBody AlgorithmRequestBody
	if err := c.BindJSON(&algorithmRequestBody); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if len(algorithmRequestBody.Algorithm) > MaxRequestBuffer {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	resp, err := h.algorithmComplexityClient.CountComplexity(c, &pb.CalculateComplexityRequest{
		Language: "javascript",
		CodeBody: algorithmRequestBody.Algorithm,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": resp.ComplexityDescription,
	})
}
