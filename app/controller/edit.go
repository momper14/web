package controller

import (
	"net/http"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/kategorien"
)

// EditController controller 1 for edit
func EditController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	type (
		Kategorie struct {
			Name string
			Sub  []string
		}

		Edit struct {
			Kategorien []Kategorie
		}
	)

	var (
		data Edit
		err  error
	)

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	kats, err := kategorien.New().AlleKategorien()
	if err != nil {
		internalError(err, w)
	}

	for _, kat := range kats {
		kat := Kategorie{
			Name: kat.ID,
			Sub:  kat.Unterkategorie,
		}

		data.Kategorien = append(data.Kategorien, kat)
	}

	customExecuteTemplate(w, r, templates.Edit, data)
}

// Edit2Controller controller 2 for edit
func Edit2Controller(w http.ResponseWriter, r *http.Request) {

	type (
		Karte struct {
			Nr    int
			Titel string
		}

		Edit2 struct {
			Name        string
			Kategorie   string
			SubKat      string
			Fortschritt int
			Anzahl      int
			Karten      []Karte
		}
	)

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
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
