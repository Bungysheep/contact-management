package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/bungysheep/contact-management/pkg/protocol/db"
	"github.com/bungysheep/contact-management/pkg/protocol/grpc"
	_ "github.com/lib/pq"
)

func main() {
	if err := runServer(); err != nil {
		log.Fatalf("ERROR: %v", err)
		os.Exit(1)
	}
}

func runServer() error {
	ctx := context.Background()

	db, err := db.OpenDbConn()
	if err != nil {
		return err
	}

	grpcServer := &grpc.Server{}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Printf("gRpc server is shutting down...\n")
			grpcServer.GetGrpcServer().GracefulStop()

			log.Printf("Listener is closing...\n")
			grpcServer.GetListener().Close()

			log.Printf("Database connection is closing...\n")
			db.Close()

			<-ctx.Done()
		}
	}()

	if err := grpcServer.RunServer(ctx, db); err != nil {
		return err
	}

	return nil
}
