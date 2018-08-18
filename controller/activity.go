package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/co3k/go-webvuln/model"
	"github.com/co3k/go-webvuln/view"
	"github.com/derekstavis/go-qs"
	"html/template"
	"net/http"
	"strconv"
)

type Activity struct {
	DB *sql.DB
}

func (a *Activity) Home(w http.ResponseWriter, r *http.Request, u model.User) {
	var params map[string]interface{}

	params, err := qs.Unmarshal(r.URL.RawQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	if params["size"] == nil {
		params["size"] = "10"
	}
	sizeInt, err := strconv.Atoi(fmt.Sprintf("%v", params["size"]))
	if err != nil {
		sizeInt = 10
	}

	if params["page"] == nil {
		params["page"] = "1"
	}
	pageInt, err := strconv.Atoi(fmt.Sprintf("%v", params["page"]))
	if err != nil {
		pageInt = 1
	}

	activities, err := model.GetActivities(a.DB, sizeInt, (pageInt-1)*sizeInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	params["activities"] = activities

	totalNum, err := model.CountActivities(a.DB)
	if err != nil {
		totalNum = 0
	}
	params["total_num"] = totalNum

	totalPage := int(totalNum/sizeInt) + 1
	params["total_page"] = totalPage
	params["sizes"] = []string{"5", "10", "20", "30", "40", "50"}
	params["user"] = u

	if pageInt > 1 {
		params["prev_pager"] = map[string]string{
			"href": fmt.Sprintf("/?page=%d&size=%d", (pageInt - 1), sizeInt),
		}
	}
	if pageInt < totalPage {
		params["next_pager"] = map[string]string{
			"href": fmt.Sprintf("/?page=%d&size=%d", (pageInt + 1), sizeInt),
		}
	}

	j, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	view.RenderHtml(w, "templates/home.tmpl", struct {
		Params template.JS
	}{
		template.JS(string(j)),
	})
}

func (a *Activity) PostActivity(w http.ResponseWriter, r *http.Request, u model.User) {
	err := model.CreateActivity(a.DB, u, r.FormValue("body"))
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
