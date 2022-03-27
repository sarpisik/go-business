package main_test

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/sarpisik/go-business/app"
	"github.com/sarpisik/go-business/config"
)

var a app.App

func TestMain(m *testing.M) {
	var err error
	p := config.Config("POSTGRES_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println(err.Error())

		panic("failed to parse the database port")
	}

	// Create a test db
	host := config.Config("POSTGRES_HOSTNAME")
	user := config.Config("POSTGRES_USER")
	password := config.Config("POSTGRES_PASSWORD")
	dbName := "test_postgres"
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s sslmode=disable",
		host,
		port,
		user,
		password,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		panic(err)
	}

	// Close db connection as the app will create a new connection on
	// the initializing stage.
	err = db.Close()
	if err != nil {
		panic(err)
	}

	a = app.App{}
	a.Initialize(
		host,
		user,
		password,
		dbName,
		port,
	)

	code := m.Run()

	os.Exit(code)
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
}

// func TestEmptyTable(t *testing.T) {
// 	clearTable()

// 	req, _ := http.NewRequest("GET", "/products", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	if body := response.Body.String(); body != "[]" {
// 		t.Errorf("Expected an empty array. Got %s", body)
// 	}
// }

// func TestGetNonExistentProduct(t *testing.T) {
// 	clearTable()

// 	req, _ := http.NewRequest("GET", "/product/11", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusNotFound, response.Code)

// 	var m map[string]string
// 	json.Unmarshal(response.Body.Bytes(), &m)
// 	if m["error"] != "Product not found" {
// 		t.Errorf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", m["error"])
// 	}
// }

// // tom: rewritten function
// func TestCreateProduct(t *testing.T) {

// 	clearTable()

// 	var jsonStr = []byte(`{"name":"test product", "price": 11.22}`)
// 	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")

// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusCreated, response.Code)

// 	var m map[string]interface{}
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	if m["name"] != "test product" {
// 		t.Errorf("Expected product name to be 'test product'. Got '%v'", m["name"])
// 	}

// 	if m["price"] != 11.22 {
// 		t.Errorf("Expected product price to be '11.22'. Got '%v'", m["price"])
// 	}

// 	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
// 	// floats, when the target is a map[string]interface{}
// 	if m["id"] != 1.0 {
// 		t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
// 	}
// }

// func TestGetProduct(t *testing.T) {
// 	clearTable()
// 	addProducts(1)

// 	req, _ := http.NewRequest("GET", "/product/1", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)
// }

// func addProducts(count int) {
// 	if count < 1 {
// 		count = 1
// 	}

// 	for i := 0; i < count; i++ {
// 		a.DB.Exec("INSERT INTO products(name, price) VALUES($1, $2)", "Product "+strconv.Itoa(i), (i+1.0)*10)
// 	}
// }

// func TestUpdateProduct(t *testing.T) {

// 	clearTable()
// 	addProducts(1)

// 	req, _ := http.NewRequest("GET", "/product/1", nil)
// 	response := executeRequest(req)
// 	var originalProduct map[string]interface{}
// 	json.Unmarshal(response.Body.Bytes(), &originalProduct)

// 	var jsonStr = []byte(`{"name":"test product - updated name", "price": 11.22}`)
// 	req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")

// 	// req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(payload))
// 	response = executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	var m map[string]interface{}
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	if m["id"] != originalProduct["id"] {
// 		t.Errorf("Expected the id to remain the same (%v). Got %v", originalProduct["id"], m["id"])
// 	}

// 	if m["name"] == originalProduct["name"] {
// 		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalProduct["name"], m["name"], m["name"])
// 	}

// 	if m["price"] == originalProduct["price"] {
// 		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", originalProduct["price"], m["price"], m["price"])
// 	}
// }

// func TestDeleteProduct(t *testing.T) {
// 	clearTable()
// 	addProducts(1)

// 	req, _ := http.NewRequest("GET", "/product/1", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	req, _ = http.NewRequest("DELETE", "/product/1", nil)
// 	response = executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	req, _ = http.NewRequest("GET", "/product/1", nil)
// 	response = executeRequest(req)
// 	checkResponseCode(t, http.StatusNotFound, response.Code)
// }
