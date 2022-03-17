package routes

import (
	"github.com/gorilla/mux"

	"github.com/sarpisik/go-business/controllers"
)

func IndexRouter(router *mux.Router) {
	router.HandleFunc("/", controllers.Index()).Methods("GET")
}
