package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Unlites/knowledge_keeper/config"
	delivery "github.com/Unlites/knowledge_keeper/internal/delivery/http"
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	_ "github.com/Unlites/knowledge_keeper/migrations"
	"github.com/Unlites/knowledge_keeper/pkg/auth"
	"github.com/Unlites/knowledge_keeper/pkg/httpserver"
	"github.com/Unlites/knowledge_keeper/pkg/logger"
	"github.com/Unlites/knowledge_keeper/pkg/metrics"
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

	if cfg.Migrations.WithDowngrade {
		if err := goose.Down(sqlDbForMigrations, cfg.Migrations.Directory); err != nil {
			log.Fatal("failed to down migrations", err)
		}
	} else {
		if err := goose.Up(sqlDbForMigrations, cfg.Migrations.Directory); err != nil {
			log.Fatal("failed to up migrations", err)
		}
	}

	tokenManager := auth.NewTokenManager(&auth.TokenManagerSettings{
		SigningKey:      cfg.Auth.SigningKey,
		AccessTokenTTL:  cfg.Auth.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.RefreshTokenTTL,
	})

	passwordHasher := auth.NewPasswordHasher(&auth.PasswordHasherSettings{
		Cost: cfg.Auth.HasherCost,
	})

	repo := repository.NewRepository(db)
	usecases := usecases.NewUsecases(repo, tokenManager, passwordHasher)

	router := delivery.NewRouter(strings.Split(cfg.HttpServer.AllowedOriginsStr, " "))

	handler := delivery.NewHandler(usecases, log, router)
	handler.InitAPI()

	httpServer := httpserver.NewHttpServer(&httpserver.Settings{
		Port:           cfg.HttpServer.Port,
		Handler:        router,
		ReadTimeout:    cfg.HttpServer.ReadTimeout,
		WriteTimeout:   cfg.HttpServer.WriteTimeout,
		MaxHeaderBytes: cfg.HttpServer.MaxHeaderBytes,
	})

	go func() {
		if err := metrics.Run(cfg.Metrics.Port); err != nil {
			log.Fatal("failed to start metrics server", err)
		}
	}()

	log.Info("Metrics server started at port " + cfg.Metrics.Port)

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
