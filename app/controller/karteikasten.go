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

type (
	// KarteikastenKasten Kasten for Karteikasten
	KarteikastenKasten struct {
		SubKat       string
		Titel        string
		Anzahl       int
		Beschreibung template.HTML
		ID           string
	}

	// KarteikastenKategorie Kategorie for Karteikasten
	KarteikastenKategorie struct {
		Name   string
		Kasten []KarteikastenKasten
	}

	// KarteikastenFilter filter
	KarteikastenFilter struct {
		Kategorien      []KarteikastenFilterKategorie
		IstKategorie    string
		IstSubKategorie string
	}

	// KarteikastenFilterKategorie Kategorien for filter
	KarteikastenFilterKategorie struct {
		Name string
		Sub  []string
	}

	// KarteikastenData Data for Karteikasten
	KarteikastenData struct {
		Kategorien         []KarteikastenKategorie
		Eingeloggt         bool
		KarteikastenFilter KarteikastenFilter
	}
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

	var (
		data KarteikastenData
		err  error
	)

	data.Eingeloggt = IstEingeloggt(w, r)

	errF(fillFilter(&data.KarteikastenFilter), w)

	kategorien, err := kategorien.New().AlleKategorien()
	errF(err, w)

	for _, kat := range kategorien {
		kaesten, err := karteikaesten.New().OeffentlicheKaestenByKategorie(kat.ID)
		errF(err, w)

		if len(kaesten) > 0 {
			kategorie := KarteikastenKategorie{
				Name: kat.ID,
			}
			for _, kas := range kaesten {
				kasten := KarteikastenKasten{
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

// KarteikastenControllerFilterKategorie controller for karteikasten with kategorie filter
func KarteikastenControllerFilterKategorie(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	var (
		data           KarteikastenData
		err            error
		oberkategorie  = mux.Vars(r)["oberkategorie"]
		unterkategorie = mux.Vars(r)["unterkategorie"]
	)

	data.Eingeloggt = IstEingeloggt(w, r)

	errF(fillFilter(&data.KarteikastenFilter), w)
	data.KarteikastenFilter.IstKategorie = oberkategorie
	data.KarteikastenFilter.IstSubKategorie = unterkategorie

	kaesten, err := karteikaesten.New().OeffentlicheKaestenByOberUnterkategorie(oberkategorie, unterkategorie)
	errF(err, w)

	kategorie := KarteikastenKategorie{
		Name: oberkategorie,
	}
	for _, kas := range kaesten {
		kasten := KarteikastenKasten{
			SubKat:       kas.Unterkategorie,
			Titel:        kas.Name,
			Anzahl:       kas.AnzahlKarten(),
			Beschreibung: template.HTML(kas.Beschreibung),
			ID:           kas.ID,
		}
		kategorie.Kasten = append(kategorie.Kasten, kasten)
	}
	data.Kategorien = append(data.Kategorien, kategorie)

	customExecuteTemplate(w, r, templates.Karteikasten, data)
}

func fillFilter(data *KarteikastenFilter) error {
	kats, err := kategorien.New().AlleKategorien()
	if err != nil {
		return err
	}

	for _, kat := range kats {
		kat := KarteikastenFilterKategorie{
			Name: kat.ID,
			Sub:  kat.Unterkategorie,
		}

		data.Kategorien = append(data.Kategorien, kat)
	}

	return nil
}
