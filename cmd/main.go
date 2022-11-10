package main

import (
	"bloggy"
	"bloggy/package/handler"
	"log"
)

func main() {
	srv := new(bloggy.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s", err.Error())
	}
}
