package grpc

import (
	"context"
	"log"
	"net"

	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	contactsystemservice "github.com/bungysheep/contact-management/pkg/service/contactsystem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server - gRpc server
type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

// GetGrpcServer - Return gRpc server
func (s *Server) GetGrpcServer() *grpc.Server {
	return s.grpcServer
}

// GetListener - Return listener
func (s *Server) GetListener() net.Listener {
	return s.listener
}

// RunServer - Run gRpc server
func (s *Server) RunServer(ctx context.Context) error {
	log.Printf("gRpc server is starting...\n")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return err
	}

	// Define server options
	opts := []grpc.ServerOption{}

	server := grpc.NewServer(opts...)

	// Register services
	contactsystemapi.RegisterContactSystemServiceServer(server, contactsystemservice.NewContactSystemService())

	// Register reflection
	reflection.Register(server)

	s.grpcServer = server
	s.listener = lis

	log.Printf("gRpc server is listening on port 50051...\n")
	return server.Serve(lis)
}
