package main

import (
	_ "hello-world-go-api/docs"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type server struct{}

// @title Hello World API
// @version 1.0
// @description This is a hello world API, with swagger annotations.
// @license.name Apache 2.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(5 * time.Second))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/hello", func(r chi.Router) {
			r.Get("/", getHello)
			r.Get("/{name}", getHelloWithName)
		})
	})

	http.ListenAndServe(":8080", r)
}
