package grpcserver

import (
	"log"
	"net"

	pb "github.com/realPointer/url-shortener/pkg/pb"
	"google.golang.org/grpc"
)

type GrpcServer interface {
	RegisterService(server pb.ShortenerServer)
	Start()
}

type Server struct {
	grpcServer *grpc.Server
}

func New() GrpcServer {
	return &Server{
		grpcServer: grpc.NewServer(),
	}
}

func (s *Server) RegisterService(server pb.ShortenerServer) {
	pb.RegisterShortenerServer(s.grpcServer, server)
}

func (s *Server) Start() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
