package view

import (
	"log"
	"net/http"
	"text/template"
)

func RenderHtml(w http.ResponseWriter, path string, params interface{}) {
	tmpl, err := template.ParseFiles(path)
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
