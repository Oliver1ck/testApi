package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Oliver1ck/testApi/internal/config"
	"github.com/Oliver1ck/testApi/internal/database"
	"github.com/Oliver1ck/testApi/internal/handler"
	"github.com/Oliver1ck/testApi/internal/repository"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига %v", err)
	}
	pool, err := database.NewPostgresPool(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}

	userRepo := repository.NewUserRepository(pool)
	handler := handler.NewAuthHandler(userRepo)
	http.HandleFunc("/register/", handler.RegisterHandler)
	// http.HandleFunc("/login/", handler.LoginHandler)
	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера %v", err)
	}
}
