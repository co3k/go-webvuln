package controller

import (
	"database/sql"
	"fmt"
	"github.com/co3k/go-webvuln/model"
	"github.com/co3k/go-webvuln/view"
	"net/http"
	"net/url"
)

type Login struct {
	DB *sql.DB
}

func (l *Login) LoginForm(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	view.RenderHtml(w, "templates/login.tmpl", struct {
		Username string
		IsError  bool
	}{
		q.Get("username"),
		len(q.Get("username")) > 0,
	})
}

func (l *Login) Authenticate(w http.ResponseWriter, r *http.Request) {
	user, err := model.UserLoginByUsernameAndPassword(l.DB, r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login?username="+url.QueryEscape(r.FormValue("username")), http.StatusFound)
		return
	}

	sessionId, err := model.GenerateSession(l.DB, user.Id)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login?username="+url.QueryEscape(r.FormValue("username")), http.StatusFound)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: sessionId,
		Path:  "/",
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func (l *Login) Deauthenticate(w http.ResponseWriter, r *http.Request, u model.User) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: "",
		Path:  "/",
	})

	err := model.ClearUserSession(l.DB, u.Id)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
