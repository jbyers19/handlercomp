package handlercomp

import (
	"net/http"

	"github.com/go-chi/chi"
)

// NewChiServer creates a new http.Server instance using Chi as the handler.
//
//	addr: address the server should run on (Ex: ":8080").
func NewChiServer(addr string) *http.Server {
	// Initialize new Chi instance.
	chiRouter := chi.NewRouter()

	// Register route.
	chiRouter.Get("/user/{id}", chiHandler)

	// Return a new http.Server using Chi as the handler.
	return &http.Server{
		Addr:    addr,
		Handler: chiRouter,
	}
}

// chiHandler handles response for the Chi server.
func chiHandler(w http.ResponseWriter, r *http.Request) {
	resString := responseBody(chi.URLParam(r, "id"), r.URL.Query().Get("name"))
	_, _ = w.Write([]byte(resString))
}
