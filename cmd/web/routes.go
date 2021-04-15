package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vitaliykolodenko/go-course/pkg/config"
	"github.com/vitaliykolodenko/go-course/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
