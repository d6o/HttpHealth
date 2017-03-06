package main

import (
	"log"
	"net/http"

	"github.com/disiqueira/itiexact"
	"github.com/disiqueira/itimethod"
	"github.com/disiqueira/itinerary"
	"github.com/disiqueira/itiwildcard"
)

// App holds the main resources from the execution.
type App struct {
	RouterPing  *itinerary.Router
	RouterAdmin *itinerary.Router
	URLList     []string
}

// Initialize receives all configs and initialize the program.
func (a *App) Initialize() {
	a.RouterPing = itinerary.NewRouter()
	a.RouterAdmin = itinerary.NewRouter()
	a.SetRoutes()
}

// Run binds the port and keeps listening for connections.
func (a *App) Run() {
	go func() {
		log.Fatal(http.ListenAndServe(":8000", a.RouterPing))
	}()
	log.Fatal(http.ListenAndServe(":8001", a.RouterAdmin))
}

// SetRoutes associates all routes with their handlers.
func (a *App) SetRoutes() {
	a.RouterPing.NewPath(a.serveList).
		AddMatcher(itiexact.New("/")).
		AddMatcher(itimethod.New("GET"))

	a.RouterAdmin.NewPath(a.serveURLList).AddMatcher(itiexact.New("/url")).
		AddMatcher(itimethod.New("GET"))
	a.RouterAdmin.NewPath(a.serveURLCreate).AddMatcher(itiexact.New("/url")).
		AddMatcher(itimethod.New("POST"))
	a.RouterAdmin.NewPath(a.serveURLDelete).AddMatcher(itiwildcard.New("/url/*")).
		AddMatcher(itimethod.New("DELETE"))
}

func (a *App) AddURL(url string) {
	a.URLList = append(a.URLList, url)
}
