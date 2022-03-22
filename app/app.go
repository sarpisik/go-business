package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/sarpisik/go-business/models"
	"github.com/sarpisik/go-business/routes"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(host, user, password, dbName string, port uint64) {
	var err error
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName,
	)

	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := a.DB.Exec(models.UserTableCreationQuery); err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	routes.SignupRouter(a.Router, a.DB)
	routes.LoginRouter(a.Router, a.DB)
	routes.LogoutRouter(a.Router)
	routes.DeleteAccountRouter(a.Router, a.DB)
	routes.UsersRouter(a.Router, a.DB)
	routes.IndexRouter(a.Router, a.DB)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
