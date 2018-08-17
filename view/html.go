package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderHtml(w http.ResponseWriter, path string, params interface{}) {
	fname := filepath.Base(path)
	tmpl, err := template.New(fname).Delims("[[", "]]").ParseFiles(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = tmpl.Execute(w, params)
	if err != nil {
		log.Fatal(err)
		return
	}
}
