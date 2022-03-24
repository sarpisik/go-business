package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarpisik/go-business/controllers"
	"github.com/sarpisik/go-business/middlewares"
)

func SignupRouter(router *mux.Router, DB *sql.DB) {
	router.HandleFunc("/signup", controllers.SignupGet()).Methods(http.MethodGet)
	router.HandleFunc("/signup", middlewares.ValidateSignupFormData(controllers.SignupPost(DB))).Methods(http.MethodPost)
}
