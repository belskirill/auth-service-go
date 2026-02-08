package app

import (
	grpcapp "proto-auth/internal/app/grpc"

	"go.uber.org/zap"
)

type App struct {
	GRPCApp *grpcapp.App
}

func NewApp(logger *zap.Logger, port int) *App {

	grpcApp := grpcapp.NewAppGRPC(logger, port)

	return &App{
		GRPCApp: grpcApp,
	}
}
