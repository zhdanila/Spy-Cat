package app

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sca/internal/config"
	db2 "sca/internal/db"
	http2 "sca/internal/delivery/http"
	"sca/internal/logger"
	"sca/internal/repository"
	"sca/internal/server"
	"sca/internal/service"
	"syscall"
	"time"
)

func Run() {
	//init config
	cfg := config.New()

	//init logger
	appLog, err := logger.NewLogger(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer appLog.Sync()

	db, err := db2.NewPostgresDB(cfg)
	if err != nil {
		appLog.Error(fmt.Sprintf("failed to initialize db: %s\n", err.Error()))
	}

	//services, repos, api handlers
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := http2.NewHandler(services, appLog)

	//init server
	srv := server.NewServer(cfg, handlers.InitRoutes())

	//running server
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			appLog.Error(fmt.Sprintf("error occurred while running http server: %s\n", err.Error()))
		}
	}()

	appLog.Info("Server started")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		appLog.Error(fmt.Sprintf("failed to stop server: %v", err))
	}
}
