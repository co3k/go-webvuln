package middleware

import (
	"database/sql"
	"github.com/co3k/go-webvuln/model"
	"net/http"
)

func Auth(db *sql.DB, f func(http.ResponseWriter, *http.Request, model.User)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		user, err := model.FindUserBySessionId(db, c.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		f(w, r, user)
	}
}
