package controller

import (
	"html/template"
	"net/http"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/Momper14/weblib/client/lernen"
)

// MeineKarteienController controller for meinekarteien
func MeineKarteienController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	type Kasten struct {
		Kategorie    string
		SubKat       string
		Titel        string
		Anzahl       int
		Beschreibung template.HTML
		Sichtbarkeit string
		Fortschritt  int
		ID           string
	}

	type Data struct {
		Selbst []Kasten
		Andere []Kasten
	}

	var (
		data   Data
		kasten Kasten
		err    error
	)
	user := GetUser()

	lernen, err := lernen.New().GelerntVonUser(user)
	if err != nil {
		internalError(err, w, r)
	}

	for _, lerne := range lernen {
		kas, err := karteikaesten.New().KastenByID(lerne.Kasten)
		if err != nil {
			internalError(err, w, r)
		}
		kasten = Kasten{
			Kategorie:    kas.Kategorie,
			SubKat:       kas.Unterkategorie,
			Titel:        kas.Name,
			Anzahl:       kas.AnzahlKarten(),
			Beschreibung: template.HTML(kas.Beschreibung),
			ID:           kas.ID,
		}
		if kasten.Fortschritt, err = kas.Fortschritt(user); err != nil {
			internalError(err, w, r)
		}

		if kas.Public {
			kasten.Sichtbarkeit = "Öffentlich"
		} else {
			kasten.Sichtbarkeit = "Privat"
		}

		if kas.Autor == user {
			data.Selbst = append(data.Selbst, kasten)
		} else {
			data.Andere = append(data.Andere, kasten)
		}
	}

	customExecuteTemplate(w, r, templates.MeineKarteien, data)
}
