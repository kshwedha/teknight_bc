package server

import (
	"net/http"

	"teknight_bc/src/api"
	"teknight_bc/src/controllers"

	"github.com/gorilla/mux"
)

func root(http.ResponseWriter, *http.Request) {
}

func users(http.ResponseWriter, *http.Request) {
}

func doesUserExists(http.ResponseWriter, *http.Request) {
}

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", api.Root).Methods("GET")
	router.HandleFunc("/users", api.Users).Methods("GET")
	router.HandleFunc("/user/{name}", api.DoesUserExists).Methods("GET")
	router.HandleFunc("/heathstatus", controllers.HealthStatus).Methods("GET")
	return router
}
