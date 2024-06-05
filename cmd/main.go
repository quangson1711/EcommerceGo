package main

import (
	"Ecommerce-Go/cmd/api"
	"log"
)

func main() {
	server := api.NewApiServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
