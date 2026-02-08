package grpcapp

import (
	"fmt"
	"net"
	"proto-auth/internal/transport/grpc/auth"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	logger     *zap.Logger
	port       int
	grpcServer *grpc.Server
}

func NewAppGRPC(
	logger *zap.Logger,
	grpcPort int,
) *App {
	GRPCServer := grpc.NewServer()

	auth.RegisterAuthAPIServer(GRPCServer)

	return &App{
		grpcServer: GRPCServer,
		logger:     logger,
		port:       grpcPort,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpc.app run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := a.grpcServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() error {
	const op = "grpc.app stop"

	a.logger.Info("stopping GRPC server",
		zap.Int("port", a.port),
		zap.String("op", op),
	)
	a.grpcServer.GracefulStop()
	return nil
}
