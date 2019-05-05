package controller

import (
	"html/template"
	"net/http"
	templates "web/template"
)

func MeineKarteienController(w http.ResponseWriter, r *http.Request) {

	const BESCHREIBUNG = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. <strong> Pellentesque risus mi </strong>, tempus quis placerat ut, porta nec nulla.Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum <a> felis venenatis </a> efficitur. Aenean ac <em>eleifend lacus </em>, in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque.Donec dui urna, vehicula et sem eget, facilisis sodales sem."

	type Kasten struct {
		Kategorie    string
		SubKat       string
		Titel        string
		Anzahl       int
		Beschreibung template.HTML
		Sichtbarkeit string
		Fortschritt  int
	}

	type Data struct {
		Selbst []Kasten
		Andere []Kasten
	}

	data := Data{
		Selbst: []Kasten{
			Kasten{
				Kategorie:    "Naturwissenschaften",
				SubKat:       "Mathematik",
				Titel:        "Geometrische Formen und Körper",
				Anzahl:       23,
				Beschreibung: template.HTML(BESCHREIBUNG),
				Sichtbarkeit: "Öffentlich",
				Fortschritt:  76,
			}, Kasten{
				Kategorie:    "Naturwissenschaften",
				SubKat:       "Chemie",
				Titel:        "Atome A-Z",
				Anzahl:       23,
				Beschreibung: template.HTML(BESCHREIBUNG),
				Sichtbarkeit: "Privat",
				Fortschritt:  20,
			}, Kasten{
				Kategorie:    "Sprachen",
				SubKat:       "Latein",
				Titel:        "Vokabeln Lekrtion 1",
				Anzahl:       23,
				Beschreibung: template.HTML(BESCHREIBUNG),
				Sichtbarkeit: "Öffentlich",
				Fortschritt:  0,
			}, Kasten{
				Kategorie:    "Gesellschaft",
				SubKat:       "Verkehrskunde",
				Titel:        "Theoriefragen Fahrprüfung",
				Anzahl:       23,
				Beschreibung: template.HTML(BESCHREIBUNG),
				Sichtbarkeit: "Öffentlich",
				Fortschritt:  6,
			},
		}, Andere: []Kasten{
			Kasten{
				Kategorie:    "Naturwissenschaften",
				SubKat:       "Physik",
				Titel:        "Lorem Ipsum",
				Anzahl:       23,
				Beschreibung: template.HTML(BESCHREIBUNG),
				Sichtbarkeit: "Öffentlich",
				Fortschritt:  100,
			},
		},
	}

	customExecuteTemplate(w, r, templates.MeineKarteien, data)
}
