package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	grpc_emp "github.com/sheelendar196/go-projects/grpc_project/internal/handler/grpc"
	"github.com/sheelendar196/go-projects/grpc_project/internal/infrastructure/repository"
	logger "golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8010"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := startServer(ctx)
	shutdown(server)
}

func startServer(ctx context.Context) *grpc_emp.Service {

	employeeRepo := repository.NewEmployeeRepository(nil)

	grpcServer := grpc.NewServer()
	grpcSrv := grpc_emp.NewService(grpcServer, employeeRepo)

	go func() {
		logger.Info("connect to http://localhost:%s/ for grpc server", port)
		if err := grpcSrv.Start(ctx, port); err != nil {
			panic("grpc server failed")
		}
	}()
	return grpcSrv
}

func shutdown(srv *grpc_emp.Service) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	defer os.Exit(0)
	defer logger.Info("shutting down grpc")
	srv.Stop()
}
