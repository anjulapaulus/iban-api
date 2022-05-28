package router

import (
	"net/http"

	"github.com/anjulapaulus/iban-api/transport/http/controllers"
	"github.com/anjulapaulus/iban-api/transport/http/middleware"
	"github.com/gorilla/mux"
)

// Init initializes the router.
func Init() *mux.Router {

	// create new router
	r := mux.NewRouter()

	requestCheckerMiddleware := middleware.NewRequestCheckerMiddleware()

	// add middleware to router
	// middleware will execute in the order they are added to the router

	// add CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(requestCheckerMiddleware.Middleware)

	// initialize controllers
	controller := controllers.NewIBANController()

	// bind controller functions to routes
	r.HandleFunc("/valid/iban", controller.CheckIBAN).Methods(http.MethodPost)

	return r
}
