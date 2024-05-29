package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	delegate *gin.Engine
}

func New() *Server {
	delegate := gin.New()
	return &Server{delegate}
}

func (s *Server) Start() error {
	s.registerRoutes()
	addr := getServerAddr()
	return s.delegate.Run(addr)
}

func getServerAddr() string {
	hostname := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return hostname + ":" + port
}
