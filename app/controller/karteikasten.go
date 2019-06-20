package controller

import (
	"html/template"
	"net/http"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/Momper14/weblib/client/kategorien"
	"github.com/Momper14/weblib/client/lernen"
	"github.com/gorilla/mux"
)

// KarteikastenControllerDelete controller for karteikasten delete
func KarteikastenControllerDelete(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	var (
		user     = GetUser(w, r)
		kaesten  = karteikaesten.New()
		kastenid = mux.Vars(r)["kastenid"]
		kasten   karteikaesten.Karteikasten
		err      error
	)

	if kasten, err = kaesten.KastenByID(kastenid); err != nil {
		internalError(err, w)
	}

	if kasten.Autor == user {
		errF(kaesten.KastenLoeschen(kastenid), w)
	} else {
		lernen := lernen.New()
		lerne, err := lernen.LerneByUserAndKasten(user, kastenid)
		errF(err, w)

		errF(lernen.LoescheLerne(lerne.ID), w)
	}

	ok(w)
}

// KarteikastenController controller for karteikasten
func KarteikastenController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	type Kasten struct {
		SubKat       string
		Titel        string
		Anzahl       int
		Beschreibung template.HTML
		ID           string
	}

	type Kategorie struct {
		Name   string
		Kasten []Kasten
	}

	type Data struct {
		Kategorien []Kategorie
		Eingeloggt bool
	}

	var (
		data      Data
		kategorie Kategorie
		kasten    Kasten
		err       error
	)

	data.Eingeloggt = IstEingeloggt(w, r)

	kategorien, err := kategorien.New().AlleKategorien()
	errF(err, w)

	for _, kat := range kategorien {
		kaesten, err := karteikaesten.New().OeffentlicheKaestenByKategorie(kat.ID)
		errF(err, w)

		if len(kaesten) > 0 {
			kategorie = Kategorie{
				Name: kat.ID,
			}
			for _, kas := range kaesten {
				kasten = Kasten{
					SubKat:       kas.Unterkategorie,
					Titel:        kas.Name,
					Anzahl:       kas.AnzahlKarten(),
					Beschreibung: template.HTML(kas.Beschreibung),
					ID:           kas.ID,
				}
				kategorie.Kasten = append(kategorie.Kasten, kasten)
			}
			data.Kategorien = append(data.Kategorien, kategorie)
		}
	}

	customExecuteTemplate(w, r, templates.Karteikasten, data)
}
