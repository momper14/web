package controller

import (
	"html/template"
	"net/http"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/gorilla/mux"
)

// ViewController controller vor view
func ViewController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

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

	var (
		err      error
		data     Data
		karte    Karte
		kastenid = mux.Vars(r)["kastenid"]
		userid   = GetUser()
	)

	kasten, err := karteikaesten.New().KastenByID(kastenid)
	if err != nil {
		internalError(err, w, r)
	}

	data = Data{
		Headline: Headline{
			Name:      kasten.Name,
			Kategorie: kasten.Kategorie,
			SubKat:    kasten.Unterkategorie,
			Ersteller: kasten.Autor,
			Anzahl:    kasten.AnzahlKarten(),
		},
	}

	if data.Headline.Fortschritt, err = kasten.Fortschritt(userid); err != nil {
		if _, ok := err.(client.NotFoundError); ok {
			data.Headline.Fortschritt = 0
		} else {
			internalError(err, w, r)
		}
	}

	if kasten.AnzahlKarten() > 0 {

		data.Titel = kasten.Karten[0].Titel
		data.Frage = template.HTML(kasten.Karten[0].Frage)
		data.Antwort = template.HTML(kasten.Karten[0].Antwort)

		fach, err := kasten.FachVonKarte(userid, 0)
		if err != nil {
			if _, ok := err.(client.NotFoundError); ok {
				fach = 0
			} else {
				internalError(err, w, r)
			}
		}

		switch fach {
		case 0:
			data.F0 = true
			break
		case 1:
			data.F1 = true
			break
		case 2:
			data.F2 = true
			break
		case 3:
			data.F3 = true
			break
		case 4:
			data.F4 = true
			break
		}

		for i, v := range kasten.Karten {
			karte = Karte{
				Nr:    i + 1,
				Titel: v.Titel,
			}
			data.Karten = append(data.Karten, karte)
		}

		data.Karten[0].Aktiv = true
	}

	customExecuteTemplate(w, r, templates.View, data)
}
