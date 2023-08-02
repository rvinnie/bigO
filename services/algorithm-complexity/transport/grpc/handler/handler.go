package handler

import (
	"context"
	"errors"
	"fmt"
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
	systemMessage := fmt.Sprintf("You will be provided with %s code, and your task is to calculate its time complexity in O-notation.", request.Language)

	if request.CodeBody == "" || request.Language == "" {
		return nil, errors.New("wrong code body or language")
	}

	resp, err := h.openAIManager.MakeRequest(request.CodeBody, systemMessage)
	if err != nil {
		return nil, err
	}

	complexityResponse := &pb.CalculateComplexityResponse{
		ComplexityDescription: resp,
	}

	return complexityResponse, nil
}
