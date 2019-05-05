package controller

import (
	"html/template"
	"net/http"
	templates "web/template"
)

func ViewController(w http.ResponseWriter, r *http.Request) {
	const BESCHREIBUNG = "<html>Lorem ipsum dolor sit amet, consectetur adipiscing elit. <strong> Pellentesque risus mi </strong>, tempus quis placerat ut, porta nec nulla.Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum <a> felis venenatis </a> efficitur. Aenean ac <em>eleifend lacus </em>, in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque.Donec dui urna, vehicula et sem eget, facilisis sodales sem.</html>"

	type Headline struct {
		Name        string
		Kategorie   string
		SubKat      string
		Ersteller   string
		Fortschritt int
		Anzahl      int
	}

	type Karte struct {
		Nr    int
		Titel string
		Aktiv bool
	}

	type Data struct {
		Headline           Headline
		Titel              string
		F0, F1, F2, F3, F4 bool
		Frage              template.HTML
		Antwort            template.HTML
		Karten             []Karte
	}

	edit := Data{
		Headline: Headline{
			Name:        "Geometrie",
			Kategorie:   "Naturwissenschaften",
			SubKat:      "Mathematik",
			Ersteller:   "Max Mustermann",
			Fortschritt: 0,
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
		Karten: []Karte{
			Karte{
				Nr:    1,
				Titel: "Titel der Karte",
				Aktiv: true,
			}, Karte{
				Nr:    2,
				Titel: "Titel der Karte",
				Aktiv: false,
			}, Karte{
				Nr:    3,
				Titel: "Titel der Karte",
				Aktiv: false,
			}, Karte{
				Nr:    4,
				Titel: "Titel der Karte",
				Aktiv: false,
			}, Karte{
				Nr:    5,
				Titel: "Titel der Karte",
				Aktiv: false,
			}, Karte{
				Nr:    6,
				Titel: "Titel der Karte",
				Aktiv: false,
			},
		},
	}

	customExecuteTemplate(w, r, templates.View, edit)
}
