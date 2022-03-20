package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarpisik/go-business/controllers"
)

func LoginRouter(router *mux.Router, DB *sql.DB) {
	router.HandleFunc("/login", controllers.LoginGet()).Methods(http.MethodGet)
	router.HandleFunc("/login", controllers.LoginPost(DB)).Methods(http.MethodPost)
}
