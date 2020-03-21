package main

import (
	"log"
	"net/http"

	"centaur/osenv"
	"centaur/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oklog/run"
	"go.uber.org/dig"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("[main] godotenv.Load() failed: %v", err)
	}
	isReleaseMode, _ := osenv.Bool("SERVER_HTTP_RUN_RELEASE")
	if isReleaseMode {
		log.Println("[main] HTTP server is running in release mode")
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	c := dig.New()
	{
		_ = c.Provide(newServer)
	}

	var g run.Group
	g.Add(func() error {
		s := takeServer(c)
		if err := s.Run(); err != nil {
			return err
		}
		return nil
	}, func(err error) {
		log.Printf("[main] Executed server failed: %v", err)
	})
	_ = g.Run()
}

func newServer() *server.Server {
	args := server.NewServerArgs{
		Router:  gin.New(),
		Addr:    osenv.String("SERVER_HTTP_ADDR"),
		Routing: routing,
	}
	return server.NewServer(args)
}

func routing(router *gin.Engine) {
	if router == nil {
		return
	}
	router.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
}

func takeServer(c *dig.Container) *server.Server {
	var s *server.Server
	err := c.Invoke(func(_s *server.Server) {
		s = _s
	})
	if err != nil {
		log.Fatalf("[main] Taking server failed: %v", err)
	}
	return s
}
