package main

import (
	"log"
	"restaurant-list-api"
	"restaurant-list-api/pkg/handler"
)

func main() {
	handler := new(handler.Handler)

	srv := new(restaurant_list_api.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
