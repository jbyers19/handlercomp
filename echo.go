package handlercomp

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewEchoServer creates a new Echo instance.
func NewEchoServer() *echo.Echo {
	// Initialize new Echo instance and set up middleware.
	e := echo.New()
	e.HideBanner = true

	// Register routes.
	e.GET("/user/:id", echoHandler)

	return e
}

// StartEchoServer starts the Echo server.
//
//	e: Echo instance to start.
//	addr: address the server should run on (Ex: ":8080").
func StartEchoServer(e *echo.Echo, addr string) {
	slog.Info("Starting Echo server...")

	// Listen for OS signals to interrupt/stop the server.
	go signalInterrupt(context.Background(), e)

	// Start the server.
	if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}

// echoHandler handles responses for Echo servers.
func echoHandler(c echo.Context) error {
	resString := responseBody(c.Param("id"), c.QueryParam("name"))
	return c.String(http.StatusOK, resString)
}
