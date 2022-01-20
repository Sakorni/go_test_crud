package router

import (
	"context"
	"net/http"
	"time"
)

type Server struct{
	server *http.Server
}

func (s *Server) Run(port string) error {
		s.server = &http.Server{
			Addr: ":" + port,
			MaxHeaderBytes: 1 << 20, // 1MB
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		return s.server.ListenAndServe()
}
func (s *Server) Shutdown(ctx context.Context) error{
	return s.server.Shutdown(ctx)
}