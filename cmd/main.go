package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/bungysheep/contact-management/pkg/logger"
	"github.com/bungysheep/contact-management/pkg/protocol/db"
	"github.com/bungysheep/contact-management/pkg/protocol/grpc"
	_ "github.com/lib/pq"
)

func main() {
	if err := runServer(); err != nil {
		logger.Log.Error(fmt.Sprintf("ERROR: %v", err))
		os.Exit(1)
	}
}

func runServer() error {
	ctx := context.Background()

	if err := logger.InitLog(); err != nil {
		return nil
	}

	db, err := db.OpenDbConn()
	if err != nil {
		return err
	}

	grpcServer := &grpc.Server{}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			logger.Log.Info("gRpc server is shutting down...")
			grpcServer.GetGrpcServer().GracefulStop()

			logger.Log.Info("Listener is closing...")
			grpcServer.GetListener().Close()

			logger.Log.Info("Database connection is closing...")
			db.Close()

			<-ctx.Done()
		}
	}()

	if err := grpcServer.RunServer(ctx, db); err != nil {
		return err
	}

	return nil
}
