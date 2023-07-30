package grpc

import (
	"fmt"
	pb "github.com/rvinnie/bigO/services/algorithm-complexity/pb"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	srv                       *grpc.Server
	algorithmComplexityServer pb.AlgorithmComplexityServer
}

func NewServer(algorithmComplexityServer pb.AlgorithmComplexityServer) *Server {
	return &Server{
		srv:                       grpc.NewServer(),
		algorithmComplexityServer: algorithmComplexityServer,
	}
}

func (s *Server) ListenAndServe(port string) error {
	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	pb.RegisterAlgorithmComplexityServer(s.srv, s.algorithmComplexityServer)

	if err = s.srv.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
