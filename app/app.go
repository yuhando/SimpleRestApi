package app

import (
	"github.com/yuhando/simpleapi/app/controller"
	"github.com/yuhando/simpleapi/app/handler"

	"github.com/gorilla/mux"
)

// App ...
type App struct {
	Router *mux.Router
}

// Initialize initializes the app with predefined configuration
func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.setRouters()
}

// setRouters sets the all required routers
func (app *App) setRouters() {
	app.Get("/", app.handleRequest(handler.GetHealtCheck))
	app.Options("/message", app.handleRequest(handler.HeartBeat))
	app.Post("/message", app.handleRequest(controller.PostMessage))
	app.Get("/message", app.handleRequest(controller.GetMessage))
}
