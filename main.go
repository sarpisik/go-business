package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	// Init the handlebars engine
	e := iris.Handlebars("./templates", ".html").Reload(true)

	// Register a helper.
	e.AddFunc("fullName", func(person map[string]string) string {
		return person["firstName"] + " " + person["lastName"]
	})

	app.RegisterView(e)

	app.Get("/", func(ctx iris.Context) {
		viewData := iris.Map{
			"author": map[string]string{"firstName": "Jean", "lastName": "Valjean"},
			"body":   "Life is difficult",
			"comments": []iris.Map{{
				"author": map[string]string{"firstName": "Marcel", "lastName": "Beliveau"},
				"body":   "LOL!",
			}},
		}

		ctx.View("index.html", viewData)
	})

	app.Listen(":8080")
}
