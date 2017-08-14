package main

import (
	"database/sql"
	"fmt"
	"github.com/co3k/go-webvuln/controller"
	"github.com/co3k/go-webvuln/middleware"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	db, err := sql.Open("sqlite3", "db/webvuln.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open database file. Please create database file by executing 'make db/webvuln.db'", err)
		os.Exit(1)
	}

	login := &controller.Login{DB: db}
	activity := &controller.Activity{DB: db}

	r.HandleFunc("/", middleware.DisableXSSFilter(middleware.Auth(db, activity.Home))).Methods("GET")
	r.HandleFunc("/activity", middleware.DisableXSSFilter(middleware.Auth(db, activity.PostActivity))).Methods("POST")
	r.HandleFunc("/login", middleware.DisableXSSFilter(login.LoginForm)).Methods("GET")
	r.HandleFunc("/login", middleware.DisableXSSFilter(login.Authenticate)).Methods("POST")
	r.HandleFunc("/logout", middleware.DisableXSSFilter(middleware.Auth(db, login.Deauthenticate))).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.NotFoundHandler = http.HandlerFunc(middleware.DisableXSSFilter(controller.NotFoundError))

	http.Handle("/", r)
	fmt.Println("Starting up the server")
	http.ListenAndServe(":8000", r)
}
