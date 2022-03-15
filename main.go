package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Listen(":8080")
}
