package routes

import (
	"database/sql"

	"github.com/gorilla/mux"

	"github.com/sarpisik/go-business/controllers"
)

func UsersRouter(router *mux.Router, DB *sql.DB) {
	router.HandleFunc("/users", controllers.GetUsers(DB))
}
