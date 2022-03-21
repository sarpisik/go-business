package routes

import (
	"database/sql"

	"github.com/gorilla/mux"

	"github.com/sarpisik/go-business/controllers"
	"github.com/sarpisik/go-business/middlewares"
)

func IndexRouter(router *mux.Router, DB *sql.DB) {
	router.HandleFunc("/", middlewares.GetCookie(middlewares.ParseJWT(middlewares.GetUserData(DB, controllers.Index)))).Methods("GET")
}
