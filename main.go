package main

import (
	"github.com/RianIhsan/go-ecommerce-app/config"
	"github.com/RianIhsan/go-ecommerce-app/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}
	api.StartServer(cfg)
}
