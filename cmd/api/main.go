package main

import "github.com/ashiqsabith123/api-gateway/pkg/di"

func main() {
	server := di.InitializeApi()

	server.Start()
}
