package middleware

import (
	"google.golang.org/grpc"

	"github.com/bungysheep/contact-management/pkg/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

// AddLoggerUnaryInterceptor add logger unary interceptor
func AddLoggerUnaryInterceptor() grpc.UnaryServerInterceptor {
	opts := []grpc_zap.Option{}

	grpc_zap.ReplaceGrpcLogger(logger.Log)

	return grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.UnaryServerInterceptor(logger.Log, opts...),
	)
}

// AddLoggerStreamInterceptor add logger stream interceptor
func AddLoggerStreamInterceptor() grpc.StreamServerInterceptor {
	opts := []grpc_zap.Option{}

	grpc_zap.ReplaceGrpcLogger(logger.Log)

	return grpc_middleware.ChainStreamServer(
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.StreamServerInterceptor(logger.Log, opts...),
	)
}
