package grpc

import (
	"context"
	"database/sql"
	"net"

	communicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	communicationmethodfieldapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	communicationmethodlabelapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodlabel"
	contactapi "github.com/bungysheep/contact-management/pkg/api/v1/contact"
	contactcommunicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	"github.com/bungysheep/contact-management/pkg/logger"
	"github.com/bungysheep/contact-management/pkg/protocol/grpc/middleware"
	communicationmethodservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethod"
	communicationmethodfieldservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodfield"
	communicationmethodlabelservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodlabel"
	contactservice "github.com/bungysheep/contact-management/pkg/service/v1/contact"
	contactcommunicationmethodservice "github.com/bungysheep/contact-management/pkg/service/v1/contactcommunicationmethod"
	contactsystemservice "github.com/bungysheep/contact-management/pkg/service/v1/contactsystem"
	communicationmethodserviceserver "github.com/bungysheep/contact-management/pkg/serviceserver/v1/communicationmethod"
	communicationmethodfieldserviceserver "github.com/bungysheep/contact-management/pkg/serviceserver/v1/communicationmethodfield"
	communicationmethodlabelserviceserver "github.com/bungysheep/contact-management/pkg/serviceserver/v1/communicationmethodlabel"
	contactserviceserver "github.com/bungysheep/contact-management/pkg/serviceserver/v1/contact"
	contactcommunicationmethodserviceserver "github.com/bungysheep/contact-management/pkg/serviceserver/v1/contactcommunicationmethod"
	contactsystemserviceserver "github.com/bungysheep/contact-management/pkg/serviceserver/v1/contactsystem"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	logger.Log.Info("gRpc server is starting...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Log.Error("Failed to listen port 50051", zap.String("err", err.Error()))
		return err
	}

	// Load ssl certificate
	creds, err := credentials.NewServerTLSFromFile("./cert/server.crt", "./cert/server.pem")
	if err != nil {
		logger.Log.Error("Failed to load Server Ssl certificate", zap.String("err", err.Error()))
		return err
	}

	// Define server options
	opts := []grpc.ServerOption{
		grpc.Creds(creds),
		middleware.AddLoggerUnaryInterceptor(logger.Log),
		middleware.AddLoggerStreamInterceptor(logger.Log),
		middleware.AddAuthenticationUnaryInterceptor(),
		middleware.AddAuthenticationStreamInterceptor(),
	}

	server := grpc.NewServer(opts...)

	// Register services
	communicationmethodapi.RegisterCommunicationMethodServiceServer(server, communicationmethodserviceserver.NewCommunicationMethodServiceServer(communicationmethodservice.NewCommunicationMethodService(db)))
	communicationmethodfieldapi.RegisterCommunicationMethodFieldServiceServer(server, communicationmethodfieldserviceserver.NewCommunicationMethodFieldServiceServer(communicationmethodfieldservice.NewCommunicationMethodFieldService(db)))
	communicationmethodlabelapi.RegisterCommunicationMethodLabelServiceServer(server, communicationmethodlabelserviceserver.NewCommunicationMethodLabelServiceServer(communicationmethodlabelservice.NewCommunicationMethodLabelService(db)))
	contactapi.RegisterContactServiceServer(server, contactserviceserver.NewContactServiceServer(contactservice.NewContactService(db)))
	contactcommunicationmethodapi.RegisterContactCommunicationMethodServiceServer(server, contactcommunicationmethodserviceserver.NewContactCommunicationMethodServiceServer(contactcommunicationmethodservice.NewContactCommunicationMethodService(db)))
	contactsystemapi.RegisterContactSystemServiceServer(server, contactsystemserviceserver.NewContactSystemServiceServer(contactsystemservice.NewContactSystemService(db)))

	// Register reflection
	reflection.Register(server)

	s.grpcServer = server
	s.listener = lis

	logger.Log.Info("gRpc server is listening on port 50051...")
	return server.Serve(lis)
}
