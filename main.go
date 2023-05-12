package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize the data-sources and other deps injections
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/test", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Test Route"))
	})

	http.ListenAndServe(":3000", r)
}
