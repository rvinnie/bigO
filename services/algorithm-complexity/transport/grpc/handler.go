package grpc

import (
	"context"
	pb "github.com/rvinnie/bigO/services/algorithm-complexity/pb"
)

type AlgorithmComplexityHandler struct {
	pb.UnimplementedAlgorithmComplexityServer
}

func NewAlgorithmComplexityHandler() *AlgorithmComplexityHandler {
	return &AlgorithmComplexityHandler{}
}

func (h *AlgorithmComplexityHandler) CountComplexity(ctx context.Context, request *pb.CalculateComplexityRequest) (*pb.CalculateComplexityResponse, error) {
	complexityResponse := &pb.CalculateComplexityResponse{
		ComplexityDescription: request.CodeBody,
	}

	return complexityResponse, nil
}
