package controller

import (
	"html/template"
	"net/http"
	templates "web/template"
)

func KarteikastenController(w http.ResponseWriter, r *http.Request) {

	type Kasten struct {
		SubKat       string
		Titel        string
		Anzahl       int
		Beschreibung template.HTML
	}

	type Kategorie struct {
		Name   string
		Kasten []Kasten
	}

	type Data struct {
		Kategorien []Kategorie
	}

	const BESCHREIBUNG = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. <strong> Pellentesque risus mi </strong>, tempus quis placerat ut, porta nec nulla.Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum <a> felis venenatis </a> efficitur. Aenean ac <em>eleifend lacus </em>, in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque.Donec dui urna, vehicula et sem eget, facilisis sodales sem."

	data := Data{
		Kategorien: []Kategorie{
			Kategorie{
				Name: "Naturwissenschaften",
				Kasten: []Kasten{
					Kasten{
						SubKat:       "Mathematik",
						Titel:        "Geometrische Formen und Körper",
						Anzahl:       23,
						Beschreibung: template.HTML(BESCHREIBUNG),
					},
					Kasten{
						SubKat:       "Chemie",
						Titel:        "Atome A-Z",
						Anzahl:       23,
						Beschreibung: template.HTML(BESCHREIBUNG),
					},
					Kasten{
						SubKat:       "Physik",
						Titel:        "Licht in Wellen und Teilchen - Modelle und Versuche",
						Anzahl:       23,
						Beschreibung: template.HTML(BESCHREIBUNG),
					},
					Kasten{
						SubKat:       "Mathematik",
						Titel:        "Das kleine 1x1",
						Anzahl:       23,
						Beschreibung: template.HTML(BESCHREIBUNG),
					},
				},
			}, Kategorie{
				Name: "Sprachen",
				Kasten: []Kasten{
					Kasten{
						SubKat:       "Latein",
						Titel:        "Vokabel Lektion 1",
						Anzahl:       23,
						Beschreibung: template.HTML(BESCHREIBUNG),
					},
					Kasten{
						SubKat:       "Englisch",
						Titel:        "Unit 2",
						Anzahl:       23,
						Beschreibung: template.HTML(BESCHREIBUNG),
					},
				},
			},
		},
	}

	customExecuteTemplate(w, r, templates.Karteikasten, data)
}
