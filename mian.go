package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nicitapa/firstProgect/internal/controller"
	"github.com/nicitapa/firstProgect/internal/repository"
	"github.com/nicitapa/firstProgect/internal/service"
	"log"
)

func main() {
	// Шаг 1. Подключение бд
	dsn := "host=localhost port=5432 user=postgres password=nicita130 dbname=onlineshop sslmode=disable"
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Шаг 2. Инициализируем слои приложения
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc)

	// Шаг 3. Запускаем http-server
	if err = ctrl.RunServer(":7779"); err != nil {
		log.Fatal(err)
	}

	// Шаг 4. Закрываем соединение с бд
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
