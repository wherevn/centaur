package main

import (
	"centaur/osenv"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("godotenv.Load() failed: %v", err)
	}
	httpServerAddr := osenv.String("HTTP_SERVER_ADDR")

	log.Printf("httpServerAddr: %v", httpServerAddr)
}
