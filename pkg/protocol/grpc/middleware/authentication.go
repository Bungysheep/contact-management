package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/bungysheep/contact-management/pkg/common/constant/contextkey"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

// AddAuthenticationUnaryInterceptor add authentication unary interceptor
func AddAuthenticationUnaryInterceptor() grpc.ServerOption {
	return grpc_middleware.WithUnaryServerChain(
		grpc_auth.UnaryServerInterceptor(auth()),
	)
}

// AddAuthenticationStreamInterceptor add authentication stream interceptor
func AddAuthenticationStreamInterceptor() grpc.ServerOption {
	return grpc_middleware.WithUnaryServerChain(
		grpc_auth.UnaryServerInterceptor(auth()),
	)
}

func auth() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		tokenInfo, err := parseToken(token)
		if err != nil {
			return nil, grpc.Errorf(codes.Unauthenticated, "Invalid auth token: %v", err)
		}

		newCtx := context.WithValue(ctx, contextkey.TokenInfo, tokenInfo)
		return newCtx, nil
	}
}

func parseToken(token string) (struct{}, error) {
	return struct{}{}, nil
}
