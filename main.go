package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nicitapa/firstProgect/internal/configs"
	"github.com/nicitapa/firstProgect/internal/controller"
	"github.com/nicitapa/firstProgect/internal/repository"
	"github.com/nicitapa/firstProgect/internal/service"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

// @title OnlineShop API
// @contact.name OnlineShop API Service
// @contact.url http://test.com
// @contact.email test@test.com
func main() {
	if err := configs.ReadSettings(); err != nil {
		log.Fatal(err)
	}

	// Шаг 1.1 Подключение бд
	dsn := fmt.Sprintf(`host=%s 
							port=%s 
							user=%s 
							password=%s 
							dbname=%s 
							sslmode=disable`,
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		os.Getenv("POSTGRES_PASSWORD"),
		configs.AppSettings.PostgresParams.Database,
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", configs.AppSettings.RedisParams.Host, configs.AppSettings.RedisParams.Port),
		DB:   configs.AppSettings.RedisParams.Database,
	})

	cache := repository.NewCache(rdb)
	repo := repository.NewRepository(db)
	svc := service.NewService(repo, cache)
	ctrl := controller.NewController(svc)

	// Шаг 3. Запускаем http-server
	if err = ctrl.RunServer(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun)); err != nil {
		log.Fatal(err)
	}

	// Шаг 4. Закрываем соединение с бд
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
