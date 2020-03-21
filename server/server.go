package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type RoutingFunc func(router *gin.Engine)

type Server struct {
	addr    string
	router  *gin.Engine
	routing RoutingFunc
}

type NewServerArgs struct {
	dig.In
	Router  *gin.Engine
	Addr    string
	Routing RoutingFunc
}

func NewServer(args NewServerArgs) *Server {
	return &Server{
		router:  args.Router,
		addr:    args.Addr,
		routing: args.Routing,
	}
}

func (_this *Server) Run() error {
	if _this.router == nil {
		return fmt.Errorf("router cannot be nil")
	}
	if _this.routing != nil {
		_this.routing(_this.router)
	}
	if err := _this.router.Run(_this.addr); err != nil {
		return err
	}
	return nil
}

func (_this *Server) Stop() {
}
