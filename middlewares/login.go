package middlewares

import (
	"context"
	"html/template"
	"net/http"

	"github.com/sarpisik/go-business/models"
	"github.com/sarpisik/go-business/utils/text"
	"github.com/sarpisik/go-business/validators"
)

func ValidateLoginFormData(next http.HandlerFunc) http.HandlerFunc {
	tmpl, _ := template.ParseFiles("/workspace/views/login.html")

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		u := &models.User{Email: r.PostFormValue("email"), Password: r.PostFormValue("password")}
		isErr, err := validators.LoginValidator(u)

		if isErr {
			tD := map[string]string{
				"title": "Login Page",
			}

			if err != nil {
				for _, err := range err {
					var errMsg string
					lowerCasedFieldName := text.GetFirstLetterLowered(err.Field())

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
					default:
						errMsg = "Invalid email."
					}

					tD[lowerCasedFieldName] = errMsg
				}
			} else {
				tD["otherFormError"] = "Something went wrong."
			}

			tmplErr := tmpl.Execute(w, tD)
			if tmplErr != nil {
				http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
			}
		} else {
			ctx := context.WithValue(r.Context(), "formData", u)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}

}
