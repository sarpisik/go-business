package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sarpisik/go-business/controllers"
	"github.com/sarpisik/go-business/middlewares"
)

func DeleteAccountRouter(router *mux.Router, DB *sql.DB) {
	router.HandleFunc("/delete-account", middlewares.GetCookie(middlewares.ParseJWT(middlewares.GetUserData(DB, controllers.DeleteUserByID(DB, middlewares.DestroyAuth))))).Methods(http.MethodGet)
}
