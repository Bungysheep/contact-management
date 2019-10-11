package grpc

import (
	"context"
	"log"
	"net"

	communicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	communicationmethodfieldapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	contactapi "github.com/bungysheep/contact-management/pkg/api/v1/contact"
	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	communicationmethodservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethod"
	communicationmethodfieldservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodfield"
	contactservice "github.com/bungysheep/contact-management/pkg/service/v1/contact"
	contactsystemservice "github.com/bungysheep/contact-management/pkg/service/v1/contactsystem"
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
	communicationmethodapi.RegisterCommunicationMethodServiceServer(server, communicationmethodservice.NewCommunicationMethodService())
	communicationmethodfieldapi.RegisterCommunicationMethodFieldServiceServer(server, communicationmethodfieldservice.NewCommunicationMethodFieldService())
	contactsystemapi.RegisterContactSystemServiceServer(server, contactsystemservice.NewContactSystemService())
	contactapi.RegisterContactServiceServer(server, contactservice.NewContactService())

	// Register reflection
	reflection.Register(server)

	s.grpcServer = server
	s.listener = lis

	log.Printf("gRpc server is listening on port 50051...\n")
	return server.Serve(lis)
}
