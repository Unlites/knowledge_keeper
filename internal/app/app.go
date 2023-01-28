package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Unlites/knowledge_keeper/config"
	delivery "github.com/Unlites/knowledge_keeper/internal/delivery/http"
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	_ "github.com/Unlites/knowledge_keeper/migrations"
	"github.com/Unlites/knowledge_keeper/pkg/httpserver"
	"github.com/Unlites/knowledge_keeper/pkg/logger"
	"github.com/Unlites/knowledge_keeper/pkg/postgres"
	"github.com/pressly/goose"
)

func Run() {
	log := logger.NewLogger()
	cfg, err := config.Init()

	if err != nil {
		log.Fatal("failed to initialize config", err)
	}

	db, err := postgres.NewPostgresDb(&postgres.Settings{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DbName:   cfg.Postgres.DbName,
		SslMode:  cfg.Postgres.SslMode,
	})
	if err != nil {
		log.Fatal("failed to connect to Postgres", err)
	}

	sqlDbForMigrations, err := db.DB()
	if err != nil {
		log.Fatal("failed to connect to Postgres for migrations", err)
	}

	if err := goose.Up(sqlDbForMigrations, "/app/migrations"); err != nil {
		log.Fatal("failed to make migrations", err)
	}

	repo := repository.NewRepository(db)
	usecases := usecases.NewUsecases(repo)

	router := delivery.NewRouter()

	handler := delivery.NewHandler(usecases, log, router)
	handler.InitAPI(router)

	httpServer := httpserver.NewHttpServer(&httpserver.Settings{
		Port:           cfg.HttpServer.Port,
		Handler:        router,
		ReadTimeout:    cfg.HttpServer.ReadTimeout,
		WriteTimeout:   cfg.HttpServer.WriteTimeout,
		MaxHeaderBytes: cfg.HttpServer.MaxHeaderBytes,
	})

	go func() {
		if err := httpServer.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("failed to start http server", err)
		}
	}()

	log.Info("Http server started at port " + cfg.HttpServer.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	log.Info("Shutdown http server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HttpServer.ShutdownTimeout)
	defer cancel()

	if err := httpServer.Stop(ctx); err != nil {
		log.Fatal("failed to stop http server", err)
	}

	log.Info("Http server stopped.")

}
