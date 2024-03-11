package core

import (
	"context"
	"os"
	"os/signal"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	grpc_emp "github.com/sheelendar196/go-projects/grpc_project/internal/handler/grpc"
	eKafka "github.com/sheelendar196/go-projects/grpc_project/internal/handler/kafka"
	"github.com/sheelendar196/go-projects/grpc_project/internal/infrastructure/repository"
	logger "golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

var producer *kafka.Producer

const (
	port  = "8010"
	topic = "employee_data"
)

func StartServer(ctx context.Context) *grpc_emp.Service {

	employeeRepo := repository.NewEmployeeRepository(nil)

	grpcServer := grpc.NewServer()
	empKafka, _ := eKafka.CreateKafkaProducer(topic)

	grpcSrv := grpc_emp.NewService(grpcServer, employeeRepo, empKafka)

	go func() {
		logger.Info("connect to http://localhost:%s/ for grpc server", port)
		if err := grpcSrv.Start(ctx, port); err != nil {
			logger.Error("error to start server: ", err)
			panic("grpc server failed")
		}
	}()
	return grpcSrv
}

func Shutdown(srv *grpc_emp.Service) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	defer os.Exit(0)
	defer logger.Info("shutting down grpc")
	srv.Stop()
}
