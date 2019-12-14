package rest

import (
	"context"
	"net/http"

	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	"github.com/bungysheep/contact-management/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server - Http server
type Server struct {
	httpServer *http.Server
}

// GetHTTPServer - Return Http server
func (svr *Server) GetHTTPServer() *http.Server {
	return svr.httpServer
}

// RunServer - Run Http server
func (svr *Server) RunServer(ctx context.Context) error {
	logger.Log.Info("Http server is starting...")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Load ssl certificate
	creds, err := credentials.NewClientTLSFromFile("./cert/ca.crt", "")
	if err != nil {
		logger.Log.Error("Failed to load Client Ssl certificate", zap.String("err", err.Error()))
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	mux := runtime.NewServeMux()

	if err := contactsystemapi.RegisterContactSystemServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts); err != nil {
		logger.Log.Error("Failed to start Http server", zap.String("err", err.Error()))
		return err
	}

	s := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	svr.httpServer = s

	logger.Log.Info("Htpp server is listening on port 3000...")
	return s.ListenAndServe()
}
