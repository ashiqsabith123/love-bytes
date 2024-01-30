package main

import (
	"fmt"
	"log"

	"github.com/ashiqsabith123/api-gateway/pkg/config"
	"github.com/ashiqsabith123/api-gateway/pkg/di"
	logs "github.com/ashiqsabith123/love-bytes-proto/log"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Eroor loading config")
	}

	//fmt.Println(config)

	err = logs.InitLogger("./pkg/logs/log.log")
	if err != nil {
		fmt.Println(err)
		logs.ErrLog.Fatalln("Error while initilizing logger", err)
	}

	server := di.InitializeApi(config)

	server.Start()
}
