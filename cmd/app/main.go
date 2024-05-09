package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hgtpcastro/go-expert-lab-cep-weather-api/configs"
	httphandlers "github.com/hgtpcastro/go-expert-lab-cep-weather-api/internal/infra/http_handlers"
)

func main() {
	config, err := configs.LoadConfig("./cmd/app/")
	if err != nil {
		log.Fatal("Error loading .env file\n", err)
		panic(err)
	}

	if config.WebServerPort == "" {
		config.WebServerPort = "8080"
	}

	http.HandleFunc("/", httphandlers.NewCepWeatherHandler(config).Handler)
	fmt.Printf("Server is running in port %s\n", config.WebServerPort)
	http.ListenAndServe(":"+config.WebServerPort, nil)
}
