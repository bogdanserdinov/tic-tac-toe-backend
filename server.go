package tictactoe_web

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string,handlers http.Handler) error{
	s.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: handlers,
		MaxHeaderBytes: 1<<20,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error{
	return s.httpServer.Shutdown(ctx)
}