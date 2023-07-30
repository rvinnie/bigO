package grpc

import (
	"context"
	"github.com/rvinnie/bigO/services/algorithm-complexity/openai_manager"
	pb "github.com/rvinnie/bigO/services/algorithm-complexity/pb"
)

type AlgorithmComplexityHandler struct {
	pb.UnimplementedAlgorithmComplexityServer
	openAIManager *openai_manager.OpenAIManager
}

func NewAlgorithmComplexityHandler(openAIManager *openai_manager.OpenAIManager) *AlgorithmComplexityHandler {
	return &AlgorithmComplexityHandler{
		openAIManager: openAIManager,
	}
}

func (h *AlgorithmComplexityHandler) CountComplexity(ctx context.Context, request *pb.CalculateComplexityRequest) (*pb.CalculateComplexityResponse, error) {

	resp, err := h.openAIManager.MakeRequest(request.Language, request.CodeBody)
	if err != nil {
		return nil, err
	}

	complexityResponse := &pb.CalculateComplexityResponse{
		ComplexityDescription: resp,
	}

	return complexityResponse, nil
}
