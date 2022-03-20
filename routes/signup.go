package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarpisik/go-business/controllers"
)

func SignupRouter(router *mux.Router, DB *sql.DB) {
	router.HandleFunc("/signup", controllers.SignupGet()).Methods(http.MethodGet)
	router.HandleFunc("/signup", controllers.SignupPost(DB)).Methods(http.MethodPost)
}
