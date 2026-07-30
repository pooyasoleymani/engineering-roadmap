package main

import (
	"context"
	"errors"
	"gostudy/internal/http_handler"
	"gostudy/internal/model"
	"gostudy/internal/repository"
	"gostudy/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func setupLogger() *log.Logger {
	logger := log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
	return logger
}

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer cancel()

	logger := setupLogger()
	logger.Println("Starting server...")

	repo := repository.NewUserRepository(
		make([]*model.User, 0),
		logger,
	)
	userService := service.NewUserService(repo, logger)
	hdl := http_handler.NewHTTPHandler(userService, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/register", hdl.RegisterHandler)
	mux.HandleFunc("/login", hdl.LoginHandler)
	mux.HandleFunc("/profile", hdl.GetProfile)

	log.Println("server start at port 8080")
	server := http.Server{
		Addr:        ":8080",
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	logger.Println("Server shutdown complete")
	<-ctx.Done()
}
