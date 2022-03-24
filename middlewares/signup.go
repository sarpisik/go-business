package middlewares

import (
	"context"
	"html/template"
	"net/http"
	"strings"

	"github.com/sarpisik/go-business/models"
	"github.com/sarpisik/go-business/validators"
)

var tmpl, _ = template.ParseFiles("views/signup.html")

func ValidateSignupFormData(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		u := models.User{Email: r.FormValue("email"), Password: r.FormValue("password")}
		sU := validators.SignUpUser{User: u, ConfirmPassword: r.FormValue("confirmPassword")}
		isErr, err := validators.UserValidator(&sU)
		if isErr {
			tD := map[string]interface{}{
				"title": "Signup Page",
			}

			if err != nil {
				for _, err := range err {
					fieldName := err.Field()
					firstLetter := fieldName[:1]
					lowerCasedFieldName := strings.ToLower(firstLetter) + fieldName[1:]
					var errMsg string

					switch lowerCasedFieldName {
					case "password":
						switch err.Tag() {
						case "min":
							errMsg = "Password must be min 6 length."
						case "max":
							errMsg = "Password must be max 20 length."
						default:
							errMsg = "Invalid password."
						}
					case "confirmPassword":
						errMsg = "Passwords not match."
					default:
						errMsg = "Invalid email."
					}

					tD[lowerCasedFieldName] = errMsg

				}
			} else {
				tD["otherInputError"] = "Something went wrong."
			}

			tmplErr := tmpl.Execute(w, tD)
			if tmplErr != nil {
				http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
			}
		} else {
			ctx := context.WithValue(r.Context(), "userFormData", &u)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
