package middlewares

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/sarpisik/go-business/config"
	"github.com/sarpisik/go-business/models"
)

func GenerateJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Minute * 60 * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Config("JWT_SECRET")))

	if err != nil {
		fmt.Printf("Something went wrong: %s", err.Error())

		return "", err
	}

	return tokenString, nil
}

func ParseJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tSecret := []byte(config.Config("JWT_SECRET"))
		tS := mux.Vars(r)["sToken"]
		if tS != "" {
			token, err := jwt.Parse(tS, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
				}

				return tSecret, nil
			})

			if err != nil {
				fmt.Errorf("Token has been expired.")
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "userID", claims["userID"])

				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println("Invalid token.")
				next.ServeHTTP(w, r)
			}

		} else {
			fmt.Println("Missing token.")
			next.ServeHTTP(w, r)
		}

	}
}

func GetCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cV, err := r.Cookie("session")
		if err != nil {
			fmt.Println("Missing cookie.")
			mux.Vars(r)["sToken"] = ""
		} else {
			mux.Vars(r)["sToken"] = cV.Value
		}

		next.ServeHTTP(w, r)
	}
}

func GetUserData(DB *sql.DB, next func(u *models.User) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if userID := r.Context().Value("userID"); userID != nil {
			u := models.User{ID: int(userID.(float64))}
			if err := u.GetUserByID(DB); err != nil {
				fmt.Printf("User not found by ID: %v\n", userID)
			}

			next(&u).ServeHTTP(w, r)

		} else {
			next(&models.User{}).ServeHTTP(w, r)
		}
	}
}

func SetAuth(userID int, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := GenerateJWT(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			cookie := http.Cookie{
				Name:     "session",
				Value:    tokenString,
				SameSite: http.SameSiteStrictMode,
				Path:     "/",
				HttpOnly: true,
				MaxAge:   86400,
			}
			http.SetCookie(w, &cookie)

			next.ServeHTTP(w, r)
		}
	}
}

func DestroyAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:     "session",
			Value:    "",
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   0,
		}
		http.SetCookie(w, &cookie)

		next.ServeHTTP(w, r)
	}
}
