package main

import (
	"context"

	"github.com/sheelendar196/go-projects/grpc_project/cmd/core"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := core.StartServer(ctx)
	core.Shutdown(server)
}
