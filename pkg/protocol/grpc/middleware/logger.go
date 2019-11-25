package middleware

import (
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
)

// AddLoggerUnaryInterceptor add logger unary interceptor
func AddLoggerUnaryInterceptor(logger *zap.Logger) grpc.ServerOption {
	opts := []grpc_zap.Option{}

	grpc_zap.ReplaceGrpcLogger(logger)

	return grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.UnaryServerInterceptor(logger, opts...),
	)
}

// AddLoggerStreamInterceptor add logger stream interceptor
func AddLoggerStreamInterceptor(logger *zap.Logger) grpc.ServerOption {
	opts := []grpc_zap.Option{}

	grpc_zap.ReplaceGrpcLogger(logger)

	return grpc_middleware.WithStreamServerChain(
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.StreamServerInterceptor(logger, opts...),
	)
}
