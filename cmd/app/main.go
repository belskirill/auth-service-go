package main

import (
	"context"
	"os"
	"os/signal"
	"proto-auth/internal/app"
	"proto-auth/internal/config"
	log "proto-auth/internal/lib/logger"
	"sync"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	cfg := config.LoadConfig()
	logger := log.SetupZapLogger(cfg.Env)

	application := app.NewApp(ctx, cfg.DataBase.DSN(), logger, cfg.GRPC.Port)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		logger.Info("starting gRPC server", zap.Int("port", cfg.GRPC.Port))
		application.GRPCApp.MustRun()
		logger.Info("gRPC server stopped")
	}()

	<-ctx.Done()

	if err := application.GRPCApp.Stop(); err != nil {
		logger.Fatal("failed to stop gRPC server", zap.Error(err))
	}

	wg.Wait()

	defer func() {
		_ = logger.Sync()
	}()

	logger.Info("shutdown complete")
}
