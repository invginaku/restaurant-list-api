package main

import (
	"log"
	restaurant_list_api "restaurant-list-api"
)

func main() {
	srv := new(restaurant_list_api.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
