package handlercomp

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
)

// StartHTTPServer starts a http.Server instance.
func StartHTTPServer(s *http.Server) {
	st := serverType(s)
	slog.Info(fmt.Sprintf("Starting %s server...", st))

	// Listen for OS signals to interrupt/stop the server.
	go signalInterrupt(context.Background(), s)

	// Start the server.
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// responseBody is a simple function that returns a string used as the body in handler functions.
func responseBody(id, name string) string {
	return strings.Join([]string{"ID: ", id, ", Name: ", name}, "")
}

func serverType(s *http.Server) string {
	switch s.Handler.(type) {
	case *chi.Mux:
		return "chi"
	case *gin.Engine:
		return "gin"
	default:
		return "net/http"
	}
}

// server exists so that the signalInterrupt() function can be used for any server with a Shutdown() method.
type server interface {
	Shutdown(context.Context) error
}

// signalInterrupt listens for OS signals to interrupt/stop the server.
func signalInterrupt(ctx context.Context, s server) {
	// Listen for OS signals to interrupt/stop the server.
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	// Trigger graceful shutdown.
	err := s.Shutdown(ctx)
	if err != nil {
		slog.Error(err.Error())
	}
}
