package handlercomp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewGinServer creates a new http.Server instance using Chi as the handler.
//
//	addr: address the server should run on (Ex: ":8080").
func NewGinServer(addr string) *http.Server {
	// Initialize new Gin instance.
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	// Register routes.
	g.GET("/user/:id", ginHandler)

	// Return a new http.Server using Gin as the handler.
	return &http.Server{
		Addr:    addr,
		Handler: g,
	}
}

// ginHandler handles responses for Gin server.
func ginHandler(c *gin.Context) {
	resString := responseBody(c.Param("id"), c.Query("name"))
	c.String(200, resString)
}
