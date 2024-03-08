package main

import (
	"context"
	"os"
	"os/signal"

	grpc_emp "github.com/sheelendar196/go-projects/grpc_project/internal/handler/grpc"
	"github.com/sheelendar196/go-projects/grpc_project/internal/infrastructure/repository"
	logger "golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

const (
	port = "8010"
)

func main() {
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
			logger.Error("error to start server: ", err)
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
