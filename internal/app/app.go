package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Unlites/knowledge_keeper/config"
	delivery "github.com/Unlites/knowledge_keeper/internal/delivery/http"
	"github.com/Unlites/knowledge_keeper/internal/infrastructure/repository"
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	"github.com/Unlites/knowledge_keeper/pkg/httpserver"
	"github.com/Unlites/knowledge_keeper/pkg/postgres"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("failed to initialize config: %s\n", err)
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
		log.Fatalf("failed to connect to Postgres: %s\n", err)
	}

	repo := repository.NewRepository(db)
	usecases := usecases.NewUsecases(repo)

	router := delivery.NewRouter()

	handler := delivery.NewHandler(usecases, router)
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
			log.Fatalf("failed to start http server: %s\n", err)
		}
	}()

	log.Println("Http server started at port " + cfg.HttpServer.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	log.Println("Shutdown http server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HttpServer.ShutdownTimeout)
	defer cancel()

	if err := httpServer.Stop(ctx); err != nil {
		log.Fatalf("failed to stop http server: %s", err)
	}

	log.Println("Http server stopped.")

}
