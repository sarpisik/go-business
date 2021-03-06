package database

import (
	"fmt"
	"strconv"

	"github.com/sarpisik/go-business/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ConnectDB() {
	var err error
	p := config.Config("POSTGRES_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println(err.Error())

		panic("failed to parse the database port")
	}

	configData := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("POSTGRES_HOSTNAME"),
		port,
		config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"),
		config.Config("POSTGRES_DB"),
	)

	DB, err = gorm.Open("postgres", configData)

	if err != nil {
		fmt.Println(err.Error())

		panic("failed to connect database")
	}

	fmt.Println("Connected to DB")
}
