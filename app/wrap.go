package app

import (
	"log"
	"net/http"
)

// Run the app on it's router
func (app *App) Run(host string) {
	log.Println("Starting development server at http://127.0.0.1" + host)
	log.Println("Quit the server with CONTROL-C.")
	log.Fatal(http.ListenAndServe(host, app.Router))
}

// Options wraps the routers for OPTIONS method
func (app *App) Options(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("OPTIONS")
}

// Get wraps the routers for GET method
func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the routers for POST method
func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

// RequestHandlerFunction ...
type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
