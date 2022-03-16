package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/kataras/iris/v12"

	"github.com/sarpisik/go-business/database"
	"github.com/sarpisik/go-business/models"
)

func main() {
	var sessionManager *scs.SessionManager

	app := iris.New()

	// Connect DB
	database.ConnectDB()
	defer database.DB.Close()

	// Initialize session manager
	models.RegisterSession(database.DB)
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Persist = true
	sessionManager.Store = postgresstore.New(database.DB.DB())

	app.Logger().SetLevel("debug")

	// Init the handlebars engine
	e := iris.Handlebars("./templates", ".html").Reload(true)

	// Register a helper.
	e.AddFunc("fullName", func(person map[string]string) string {
		return person["firstName"] + " " + person["lastName"]
	})

	app.RegisterView(e)

	app.UseGlobal(func(ctx iris.Context) {
		c, err := sessionManager.Load(ctx.Request().Context(), ctx.GetCookie(sessionManager.Cookie.Name))
		if err != nil {
			ctx.Redirect("/login", iris.StatusTemporaryRedirect)
		} else {
			userID := sessionManager.GetString(c, "userID")
			if userID == "" && ctx.GetCurrentRoute().Path() != "/login" {
				ctx.Redirect("/login", iris.StatusTemporaryRedirect)
			} else {
				ctx.Values().Set("userID", userID)

				ctx.Next()
			}
		}
	})

	app.Get("/login", func(ctx iris.Context) {
		ctx.View("login.html")
	})

	app.Get("/", func(ctx iris.Context) {
		userID := ctx.Values().GetString("userID")
		viewData := iris.Map{
			"author": map[string]string{"firstName": "Jean", "lastName": userID},
			"body":   "Life is difficult",
			"comments": []iris.Map{{
				"author": map[string]string{"firstName": "Marcel", "lastName": "Beliveau"},
				"body":   "LOL!",
			}},
		}

		ctx.View("index.html", viewData)
	})

	app.DoneGlobal(func(irisCtx iris.Context) {
		sessCtx := irisCtx.Request().Context()

		if sessionManager.Status(sessCtx) != scs.Unmodified {
			responseCookie := &http.Cookie{
				Name:     sessionManager.Cookie.Name,
				Path:     sessionManager.Cookie.Path,
				Domain:   sessionManager.Cookie.Domain,
				Secure:   sessionManager.Cookie.Secure,
				HttpOnly: sessionManager.Cookie.HttpOnly,
				SameSite: sessionManager.Cookie.SameSite,
			}

			switch sessionManager.Status(sessCtx) {
			case scs.Modified:
				token, _, err := sessionManager.Commit(sessCtx)
				if err != nil {
					panic(err)
				}

				responseCookie.Value = token

			case scs.Destroyed:
				responseCookie.Expires = time.Unix(1, 0)
				responseCookie.MaxAge = -1
			}

			irisCtx.SetCookie(responseCookie)
			addHeaderIfMissing(irisCtx.ResponseWriter(), "Cache-Control", `no-cache="Set-Cookie"`)
			addHeaderIfMissing(irisCtx.ResponseWriter(), "Vary", "Cookie")

		}
	})

	app.SetExecutionRules(iris.ExecutionRules{
		Done: iris.ExecutionOptions{Force: true},
	})

	app.Listen(":8080")
}

func addHeaderIfMissing(w http.ResponseWriter, key, value string) {
	for _, h := range w.Header()[key] {
		if h == value {
			return
		}
	}
	w.Header().Add(key, value)
}
