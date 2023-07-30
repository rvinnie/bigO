package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rvinnie/bigO/services/gateway/config"
	pb "github.com/rvinnie/bigO/services/gateway/pb"
	"google.golang.org/grpc"
	"net/http"
)

type Handler struct {
	algorithmComplexityClient pb.AlgorithmComplexityClient
}

func NewHandler(grpcConn grpc.ClientConnInterface) *Handler {
	return &Handler{
		algorithmComplexityClient: pb.NewAlgorithmComplexityClient(grpcConn),
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.GIN.Mode)

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type"},
	}))
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
