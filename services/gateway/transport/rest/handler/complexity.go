package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/rvinnie/bigO/services/gateway/pb"
	"net/http"
	"strings"
)

const (
	MaxRequestBody = 4096
)

func (h *Handler) initComplexityRoutes(api *gin.RouterGroup) {
	complexity := api.Group("/complexity")
	{
		complexity.POST("/count", h.countAlgorithmComplexity)
	}
}

type AlgorithmComplexityRequestBody struct {
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
}

func (h *Handler) countAlgorithmComplexity(c *gin.Context) {
	var algorithmRequestBody AlgorithmComplexityRequestBody
	if err := c.BindJSON(&algorithmRequestBody); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if len(algorithmRequestBody.Code) > MaxRequestBody {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	resp, err := h.algorithmComplexityClient.CountComplexity(c, &pb.CalculateComplexityRequest{
		Language: algorithmRequestBody.Language,
		CodeBody: algorithmRequestBody.Code,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	shortDesc := parseAlgorithmComplexity(resp.ComplexityDescription)
	fullDesc := resp.ComplexityDescription

	c.JSON(http.StatusOK, map[string]interface{}{
		"shortDescription": shortDesc,
		"fullDescription":  fullDesc,
	})
}

func parseAlgorithmComplexity(text string) string {
	startIdx := strings.Index(text, "O(")
	endIdx := -1

	if startIdx == -1 {
		return ""
	}

	for i := startIdx; i < len(text); i++ {
		if text[i] == ')' {
			endIdx = i
			break
		}
	}

	if endIdx == -1 {
		return ""
	}

	return text[startIdx : endIdx+1]
}
