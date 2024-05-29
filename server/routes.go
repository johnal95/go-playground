package server

import "github.com/johnal95/go-playground/api"

func (s *Server) registerRoutes() {
	s.delegate.GET("/api/v1/example", api.GetExampleV1)
}
