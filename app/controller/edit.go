package controller

import (
	"net/http"
	templates "web/template"
)

func EditController(w http.ResponseWriter, r *http.Request) {

	type Kategorie struct {
		Name string
		Sub  []string
	}

	type Kategorien []Kategorie

	type Edit struct {
		Kategorien Kategorien
	}

	edit := Edit{
		Kategorien: Kategorien{
			Kategorie{
				Name: "Naturwissenschaften",
				Sub:  []string{"Biologie", "Chemie", "Elektrotechnik", "Informatik", "Mathematik", "Medizin", "Naturkunde", "Physik", "Sonstiges"},
			},
			Kategorie{
				Name: "Sprachen",
				Sub:  []string{"Chinesisch", "Deutsch", "Englisch", "Französisch", "Griechisch", "Italienisch", "Latein", "Russisch", "Sonstiges"},
			},
		},
	}

	customExecuteTemplate(w, r, templates.Edit, edit)
}

func Edit2Controller(w http.ResponseWriter, r *http.Request) {

	type Karte struct {
		Nr    int
		Titel string
	}

	type Edit2 struct {
		Name        string
		Kategorie   string
		SubKat      string
		Fortschritt int
		Anzahl      int
		Karten      []Karte
	}

	data := Edit2{
		Name:        "Geometrie",
		Kategorie:   "Naturwissenschaften",
		SubKat:      "Mathematik",
		Fortschritt: 0,
		Anzahl:      23,
		Karten: []Karte{
			Karte{
				Nr:    0,
				Titel: "Titel der Karte",
			},
			Karte{
				Nr:    1,
				Titel: "Titel der Karte",
			},
			Karte{
				Nr:    2,
				Titel: "Titel der Karte",
			},
			Karte{
				Nr:    3,
				Titel: "Titel der Karte",
			},
			Karte{
				Nr:    4,
				Titel: "Titel der Karte",
			},
			Karte{
				Nr:    5,
				Titel: "Titel der Karte",
			},
		},
	}

	customExecuteTemplate(w, r, templates.Edit2, data)
}
