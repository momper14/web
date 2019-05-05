package controller

import (
	"net/http"
	templates "web/template"
)

func ProfilController(w http.ResponseWriter, r *http.Request) {

	type Data struct {
		Bild     string
		Name     string
		Email    string
		Karten   int
		Karteien int
		Seit     string
	}

	data := Data{
		Bild:     "static/res/Icons/Mein-Profil.svg",
		Name:     "Max Mustermann",
		Email:    "mustermann@example.com",
		Karten:   24,
		Karteien: 7,
		Seit:     "24.12.2018",
	}

	customExecuteTemplate(w, r, templates.Profil, data)
}
