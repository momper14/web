package controller

import (
	"html/template"
	"net/http"
	templates "web/template"
)

func LernController(w http.ResponseWriter, r *http.Request) {

	const BESCHREIBUNG = "<html>Lorem ipsum dolor sit amet, consectetur adipiscing elit. <strong> Pellentesque risus mi </strong>, tempus quis placerat ut, porta nec nulla.Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum <a> felis venenatis </a> efficitur. Aenean ac <em>eleifend lacus </em>, in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque.Donec dui urna, vehicula et sem eget, facilisis sodales sem.</html>"

	type Headline struct {
		Name               string
		Kategorie          string
		SubKat             string
		Fortschritt        int
		A0, A1, A2, A3, A4 int
		Anzahl             int
	}

	type Data struct {
		Headline           Headline
		Titel              string
		F0, F1, F2, F3, F4 bool
		Frage, Antwort     template.HTML
	}

	data := Data{
		Headline: Headline{
			Name:        "Geometrie",
			Kategorie:   "Naturwissenschaften",
			SubKat:      "Mathematik",
			Fortschritt: 0,
			A0:          5,
			A1:          6,
			A2:          2,
			A3:          1,
			A4:          0,
			Anzahl:      23,
		},
		Titel:   "Titel der Karte",
		F0:      false,
		F1:      true,
		F2:      false,
		F3:      false,
		F4:      false,
		Frage:   template.HTML(BESCHREIBUNG),
		Antwort: template.HTML(BESCHREIBUNG),
	}

	customExecuteTemplate(w, r, templates.Lern, data)
}
