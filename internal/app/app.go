package app

import (
	"context"
	grpcapp "proto-auth/internal/app/grpc"
	"proto-auth/internal/infra/repositories/postgres"
	postgresRepository "proto-auth/internal/infra/repositories/postgres/auth"
	"proto-auth/internal/usecase/auth"

	"go.uber.org/zap"
)

type App struct {
	GRPCApp *grpcapp.App
}

func NewApp(ctx context.Context, dsn string, logger *zap.Logger, port int) *App {

	db := postgres.MustConnect(ctx, dsn, logger)
	repoPostgres := postgresRepository.NewRepository(db)

	authService := auth.NewService(logger, repoPostgres, repoPostgres)

	grpcApp := grpcapp.NewAppGRPC(logger, port, authService)

	return &App{
		GRPCApp: grpcApp,
	}
}
