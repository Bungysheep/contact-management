package grpc

import (
	"context"
	"database/sql"
	"log"
	"net"

	communicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	communicationmethodfieldapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	contactapi "github.com/bungysheep/contact-management/pkg/api/v1/contact"
	contactcommunicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	communicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod"
	communicationmethodfieldrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodfield"
	contactcommunicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod"
	contactrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contact"
	contactsystemrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactsystem"
	communicationmethodservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethod"
	communicationmethodfieldservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodfield"
	contactservice "github.com/bungysheep/contact-management/pkg/service/v1/contact"
	contactcommunicationmethodservice "github.com/bungysheep/contact-management/pkg/service/v1/contactcommunicationmethod"
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
func (s *Server) RunServer(ctx context.Context, db *sql.DB) error {
	log.Printf("gRpc server is starting...\n")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return err
	}

	// Define server options
	opts := []grpc.ServerOption{}

	server := grpc.NewServer(opts...)

	// Register services
	communicationmethodapi.RegisterCommunicationMethodServiceServer(server, communicationmethodservice.NewCommunicationMethodService(communicationmethodrepository.NewCommunicationMethodRepository(db)))
	communicationmethodfieldapi.RegisterCommunicationMethodFieldServiceServer(server, communicationmethodfieldservice.NewCommunicationMethodFieldService(communicationmethodfieldrepository.NewCommunicationMethodFieldRepository(db)))
	contactsystemapi.RegisterContactSystemServiceServer(server, contactsystemservice.NewContactSystemService(contactsystemrepository.NewContactSystemRepository(db)))
	contactapi.RegisterContactServiceServer(server, contactservice.NewContactService(contactrepository.NewContactRepository(db)))
	contactcommunicationmethodapi.RegisterContactCommunicationMethodServiceServer(server, contactcommunicationmethodservice.NewContactCommunicationMethodService(contactcommunicationmethodrepository.NewContactCommunicationMethodRepository(db)))

	// Register reflection
	reflection.Register(server)

	s.grpcServer = server
	s.listener = lis

	log.Printf("gRpc server is listening on port 50051...\n")
	return server.Serve(lis)
}
