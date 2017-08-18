package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/co3k/go-webvuln/model"
	"github.com/co3k/go-webvuln/view"
	"net/http"
	"strconv"
)

type Activity struct {
	DB *sql.DB
}

func (a *Activity) Home(w http.ResponseWriter, r *http.Request, u model.User) {
	q := r.URL.Query()

	size := q.Get("size")
	if size == "" {
		size = "10"
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		sizeInt = 10
	}

	page := q.Get("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	activities, err := model.GetActivities(a.DB, sizeInt, (pageInt-1)*sizeInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	json_activities, err := json.Marshal(activities)
	if err != nil {
		fmt.Println(err)
		return
	}
	totalNum, err := model.CountActivities(a.DB)
	if err != nil {
		totalNum = 0
	}

	view.RenderHtml(w, "templates/home.tmpl", struct {
		User       model.User
		Size       string
		Page       string
		TotalNum   string
		Activities string
		Sizes      []string
	}{
		u,
		size,
		page,
		strconv.Itoa(totalNum),
		string(json_activities),
		[]string{"5", "10", "20", "30", "40", "50"},
	})
}

func (a *Activity) PostActivity(w http.ResponseWriter, r *http.Request, u model.User) {
	err := model.CreateActivity(a.DB, u, r.FormValue("body"))
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
