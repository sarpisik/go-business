package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarpisik/go-business/controllers"
)

func LogoutRouter(router *mux.Router) {
	router.HandleFunc("/logout", controllers.LogoutGet).Methods(http.MethodGet)
}
