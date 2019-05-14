package controller

import (
	"html/template"
	"net/http"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/Momper14/weblib/client/kategorien"
)

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
	}

	var (
		data      Data
		kategorie Kategorie
		kasten    Kasten
		err       error
	)

	kategorien, err := kategorien.New().AlleKategorien()
	if err != nil {
		internalError(err, w)
	}

	for _, kat := range kategorien {
		kaesten, err := karteikaesten.New().OeffentlicheKaestenByKategorie(kat.ID)
		if err != nil {
			internalError(err, w)
		}

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
