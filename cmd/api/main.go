package main

import (
	"fmt"
	"log"

	"github.com/ashiqsabith123/api-gateway/pkg/config"
	"github.com/ashiqsabith123/api-gateway/pkg/di"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Eroor loading config")
	}

	fmt.Println(config)

	server := di.InitializeApi(config)

	server.Start()
}
