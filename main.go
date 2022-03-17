package main

import (
	"fmt"
	"strconv"

	"github.com/sarpisik/go-business/app"
	"github.com/sarpisik/go-business/config"
)

func main() {
	var err error
	p := config.Config("POSTGRES_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println(err.Error())

		panic("failed to parse the database port")
	}
	a := app.App{}
	a.Initialize(
		config.Config("POSTGRES_HOSTNAME"),
		config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"),
		config.Config("POSTGRES_DB"),
		port,
	)

	a.Run(":8080")
}
