package grpc

import (
	"context"
	"database/sql"
	"log"
	"net"

	communicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	communicationmethodfieldapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	communicationmethodlabelapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodlabel"
	contactapi "github.com/bungysheep/contact-management/pkg/api/v1/contact"
	contactcommunicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	communicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod"
	communicationmethodfieldrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodfield"
	communicationmethodlabelrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodlabel"
	contactrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contact"
	contactcommunicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod"
	communicationmethodservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethod"
	communicationmethodfieldservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodfield"
	communicationmethodlabelservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodlabel"
	contactservice "github.com/bungysheep/contact-management/pkg/service/v1/contact"
	contactcommunicationmethodservice "github.com/bungysheep/contact-management/pkg/service/v1/contactcommunicationmethod"
	contactsystemservice "github.com/bungysheep/contact-management/pkg/service/v1/contactsystem"
	contactsystemserviceserver "github.com/bungysheep/contact-management/pkg/serviceserver/v1/contactsystem"
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
	communicationmethodlabelapi.RegisterCommunicationMethodLabelServiceServer(server, communicationmethodlabelservice.NewCommunicationMethodLabelService(communicationmethodlabelrepository.NewCommunicationMethodLabelRepository(db)))
	contactsystemapi.RegisterContactSystemServiceServer(server, contactsystemserviceserver.NewContactSystemServiceServer(contactsystemservice.NewContactSystemService(db)))
	contactapi.RegisterContactServiceServer(server, contactservice.NewContactService(contactrepository.NewContactRepository(db)))
	contactcommunicationmethodapi.RegisterContactCommunicationMethodServiceServer(server, contactcommunicationmethodservice.NewContactCommunicationMethodService(contactcommunicationmethodrepository.NewContactCommunicationMethodRepository(db)))

	// Register reflection
	reflection.Register(server)

	s.grpcServer = server
	s.listener = lis

	log.Printf("gRpc server is listening on port 50051...\n")
	return server.Serve(lis)
}
