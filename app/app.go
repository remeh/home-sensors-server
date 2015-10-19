// Copyright © 2015 - Rémy MATHIEU

package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Config Config
	router *mux.Router
}

func NewApp() *App {
	app := &App{}
	app.router = mux.NewRouter()
	app.Config = ReadConfig()
	return app
}

func (a *App) Add(pattern string, handler http.Handler, methods ...string) {
	a.router.PathPrefix("/").Subrouter().Handle(pattern, handler).Methods(methods...)
}

func (a *App) Start() {
	// Prepares the router serving the static pages and assets.
	a.prepareStaticRouter()

	// Handles static routes
	http.Handle("/", a.router)

	// Starts listening.
	log.Println("Start listening on:", a.Config.Addr)
	if err := http.ListenAndServe(a.Config.Addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func (a *App) prepareStaticRouter() {
	// Add the final route, the static assets and pages.
	a.router.PathPrefix("/").Handler(http.FileServer(http.Dir(a.Config.PublicDir)))
	log.Println("Serving static from directory:", a.Config.PublicDir)
}
