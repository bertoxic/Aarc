package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/bertoxic/aarc/config"
	"github.com/bertoxic/aarc/handlers"
)

func routes(app *config.AppConfig) http.Handler{

mux:= chi.NewRouter()
mux.Use(SessionsLoad)
mux.Post("/register", handlers.Repo.Register)
mux.Get("/verify", handlers.Repo.Verify)
mux.Post("/verify", handlers.Repo.PostVerify)
mux.Get("/arcform", handlers.Repo.HomePage)
mux.Post("/arcform", handlers.Repo.PostHome)
mux.Get("/recommendation", handlers.Repo.RecommendationPage)

fileServer := http.FileServer(http.Dir("./static/"))
mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
return mux
}