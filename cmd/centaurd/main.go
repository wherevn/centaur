package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("godotenv.Load() failed: %v", err)
    }
    httpServerAddr := os.Getenv("HTTP_SERVER_ADDR")

    log.Printf("httpServerAddr: %v", httpServerAddr)
}
