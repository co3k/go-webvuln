package controller

import (
	"github.com/co3k/go-webvuln/view"
	"net/http"
)

func NotFoundError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	q := r.URL.Query()
	view.RenderHtml(w, "templates/404.tmpl", struct {
		ReturnPage string
	}{
		q.Get("return_page"),
	})
}
