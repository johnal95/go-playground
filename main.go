package main

import (
	"log"

	"github.com/johnal95/go-playground/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := server.New()
	log.Fatal(server.Start())
}
