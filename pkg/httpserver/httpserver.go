package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// Config is configuration for http server.
type Config struct {
	SrvPort string `envconfig:"SERVER_PORT" default:"8080"`
	GinMode string `envconfig:"GIN_MODE" default:"debug"`
}

// Server ...
type Server struct {
	Addr           string
	Router         *gin.Engine
	OnShutdownHook func()

	server *http.Server
}

// NewServer create a server with default configuration.
func NewServer() *Server {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return NewServerWithConfig(cfg)
}

// NewServerWithConfig ...
func NewServerWithConfig(cfg Config) *Server {
	gin.SetMode(cfg.GinMode)
	return &Server{
		Addr:   ":" + cfg.SrvPort,
		Router: gin.New(),
	}
}

// Run a server on Addr.
func (srv *Server) Run() {
	srv.server = &http.Server{
		Addr:    srv.Addr,
		Handler: srv.Router,
	}

	go func(server *http.Server) {
		logrus.Infof("Listen on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("Error while listen and serve: %v", err)
		}
	}(srv.server)

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)
	<-wait

	logrus.Info("Shutting down http server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if err := srv.server.Shutdown(ctx); err != nil {
		logrus.Errorf("Cannot shutdown server: %v", err)
	}

	cancel()
	if srv.OnShutdownHook != nil {
		srv.OnShutdownHook()
	}
}
