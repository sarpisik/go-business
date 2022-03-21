package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarpisik/go-business/controllers"
	"github.com/sarpisik/go-business/middlewares"
)

func LogoutRouter(router *mux.Router) {
	router.HandleFunc("/logout", middlewares.DestroyAuth(controllers.LogoutGet)).Methods(http.MethodGet)
}
