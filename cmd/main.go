package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/benleem/show-beam/internal/handlers"
	"github.com/joho/godotenv"
)

type Config struct {
	port string
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load(".env")
	port, portOk := os.LookupEnv("PORT")
	if !portOk {
		return nil, fmt.Errorf("port not defined")
	}
	port = fmt.Sprintf(":%v", port)
	return &Config{
		port,
	}, nil
}

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	mux := handlers.Init()
	log.Fatalln(http.ListenAndServe(config.port, mux))
}
