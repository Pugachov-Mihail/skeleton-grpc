package api

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
)

type App struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	return &App{log, grpcServer, port}
}
