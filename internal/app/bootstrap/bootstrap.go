package bootstrap

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sca/internal/app/transport/http/handler"
	"sca/internal/config"
	"sca/internal/repository"
	"sca/internal/service"
	"sca/pkg/db"
	"sca/pkg/logger"
	"sca/pkg/server"
	"syscall"
)

func Website() {
	logger.InitLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := db.NewPostgresDB(cfg)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("error with connecting to database: %s", err.Error()))
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.HTTPPort, handlers.InitRoutes()); err != nil {
			zap.L().Fatal(fmt.Sprintf("error with running server: %s", err.Error()))
		}
	}()

	zap.L().Info("Spy-Cats Backend started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	zap.L().Info("Spy-Cats Backend Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		zap.L().Fatal(fmt.Sprintf("error with shutting down server: %s", err.Error()))
	}

	if err := db.Close(); err != nil {
		zap.L().Fatal(fmt.Sprintf("error with closing db: %s", err.Error()))
	}
}
