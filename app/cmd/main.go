package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"practice.dev/infra/handler"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"practice.dev/api"
)

func main() {
	ctx := context.Background()
	if err := os.Setenv("SPANNER_EMULATOR_HOST", "localhost:9010"); err != nil {
		panic(err)
	}
	d := "projects/test-project/instances/test-instance/databases/test-database"
	cli, err := spanner.NewClient(ctx, d)
	if err != nil {
		panic(err)
	}
	port := "9000"
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	server := handler.NewServer(cli)
	api.RegisterActionServiceServer(s, server)
	reflection.Register(s)

	errCh := make(chan error)
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, syscall.SIGTERM)
	go func() {
		log.Default().Println("start gRPC")
		if err := s.Serve(listener); err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-interruptCh:
		s.GracefulStop()
	}
}
