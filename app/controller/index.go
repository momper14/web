package controller

import (
	"net/http"
	templates "web/template"
)

func IndexController(w http.ResponseWriter, r *http.Request) {

	type Index struct {
		Nutzer, Karten, Karteien int
	}

	data := Index{
		Nutzer:   32,
		Karten:   124,
		Karteien: 22,
	}

	customExecuteTemplate(w, r, templates.Index, data)
}
