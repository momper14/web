package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

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
		KastenID           string
	}

	var (
		err      error
		data     Data
		kastenid = mux.Vars(r)["kastenid"]
		userid   string
		index    int
	)

	if IstEingeloggt(w, r) {
		userid = GetUser(w, r)
	}

	kasten, err := karteikaesten.New().KastenByID(kastenid)
	if err != nil {
		if _, ok := err.(client.NotFoundError); ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		internalError(err, w)
	}

	if tmp, ok := mux.Vars(r)["karte"]; ok {
		if index, err = strconv.Atoi(tmp); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if index < 0 || index >= kasten.AnzahlKarten() {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	} else {
		index = 0
	}
	fmt.Println(index)

	data = Data{
		Headline: Headline{
			Name:      kasten.Name,
			Kategorie: kasten.Kategorie,
			SubKat:    kasten.Unterkategorie,
			Ersteller: kasten.Autor,
			Anzahl:    kasten.AnzahlKarten(),
		},
		KastenID: kasten.ID,
	}

	if userid == "" {
		data.Headline.Fortschritt = 0
	} else {
		if data.Headline.Fortschritt, err = kasten.Fortschritt(userid); err != nil {
			internalError(err, w)
		}
	}

	if kasten.AnzahlKarten() > 0 {

		karte := kasten.Karten[index]
		data.Titel = karte.Titel
		data.Frage = template.HTML(karte.Frage)
		data.Antwort = template.HTML(karte.Antwort)

		var fach int
		if userid == "" {
			fach = 0
		} else {
			fach, err = kasten.FachVonKarte(userid, index)
			if err != nil {
				internalError(err, w)
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
			karte := Karte{
				Nr:    i + 1,
				Titel: v.Titel,
			}
			data.Karten = append(data.Karten, karte)
		}

		data.Karten[index].Aktiv = true
	}

	customExecuteTemplate(w, r, templates.View, data)
}
