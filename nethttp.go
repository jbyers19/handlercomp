// Package handlercomp implements several popular web frameworks based on net/http
// to compare their performance.
package handlercomp

import (
	"net/http"
)

// NewHTTPServer creates a new http.Server instance using DefaultServerMux as the handler.
//
//	addr: address the server should run on (Ex: ":8080").
func NewHTTPServer(addr string) *http.Server {
	// Use default server mutex as the handler.
	mux := http.NewServeMux()

	// Register route.
	mux.HandleFunc(http.MethodGet+" /user/{id}", httpHandler)

	// Return a new http.Server.
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

// httpHandler handles responses for the net/http server.
func httpHandler(w http.ResponseWriter, r *http.Request) {
	resString := responseBody(r.PathValue("id"), r.URL.Query().Get("name"))
	_, _ = w.Write([]byte(resString))
}
