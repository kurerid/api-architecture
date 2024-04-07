package server

import (
	"context"
	"github.com/execaus/exloggo"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
	}
	exloggo.Info("server started successfully")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(postgres *sqlx.DB, ctx context.Context) {
	exloggo.Info("server shutdown process started")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		exloggo.Error(err.Error())
	} else {
		exloggo.Info("http listener shutdown successfully")
	}

	if err := postgres.Close(); err != nil {
		exloggo.Error(err.Error())
	} else {
		exloggo.Info("business database connection closed successfully")
	}

	exloggo.Info("server shutdown process completed successfully")
}
