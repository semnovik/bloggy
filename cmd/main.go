package main

import (
	"bloggy"
	"bloggy/package/handler"
	"bloggy/package/repository"
	"bloggy/package/service"
	"context"
	"github.com/demimurg/twitter/pkg/log"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
)

func main() {
	log.SetLevel("DEBUG")

	ctx := context.Background()

	if err := InitConfig(); err != nil {
		log.Fatal(ctx, "error with reading config", "error", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(ctx, "error with reading env-variables", "error", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatal(ctx, "error while initialize db", "error", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(bloggy.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal(ctx, "error occurred while running http server", "error", err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
