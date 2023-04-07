package controllers

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
