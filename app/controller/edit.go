package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Momper14/web/app/model"
	"github.com/gorilla/mux"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/Momper14/weblib/client/kategorien"
)

type (
	// EditKategorie Kategorie for Edit
	EditKategorie struct {
		Name string
		Sub  []string
	}

	// Edit Data for Edit
	Edit struct {
		Titel           string
		Kategorien      []EditKategorie
		Beschreibung    string
		Public          bool
		IstKategorie    string
		IstSubKategorie string
	}
)

// EditControllerPut controller 1 for edit put
func EditControllerPut(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	var (
		username = GetUser(w, r)
		kaesten  = karteikaesten.New()
		err      error
		decoder  = json.NewDecoder(r.Body)
		neu      model.Karteikasten
		kasten   karteikaesten.Karteikasten
		kastenid = mux.Vars(r)["kastenid"]
	)

	if err = decoder.Decode(&neu); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		panic(err)
	}

	if neu.Beschreibung == "" || neu.Kategorie == "" || neu.Titel == "" || neu.Unterkategorie == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if kasten, err = kaesten.KastenByID(kastenid); err != nil {
		internalError(err, w)
	}

	if kasten.Autor != username {
		forbidden(w)
		return
	}

	if kasten.Name != neu.Titel || kasten.Kategorie != neu.Kategorie ||
		kasten.Unterkategorie != neu.Unterkategorie || kasten.Public != neu.Public ||
		kasten.Beschreibung != neu.Beschreibung {

		kasten.Name = neu.Titel
		kasten.Kategorie = neu.Kategorie
		kasten.Unterkategorie = neu.Unterkategorie
		kasten.Beschreibung = neu.Beschreibung
		kasten.Public = neu.Public

		if err = kaesten.KastenBearbeiten(kasten); err != nil {
			internalError(err, w)
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

// EditControllerPost controller 1 for edit post
func EditControllerPost(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	var (
		username = GetUser(w, r)
		uuid     string
		kaesten  = karteikaesten.New()
		err      error
		decoder  = json.NewDecoder(r.Body)
		neu      model.Karteikasten
		kasten   karteikaesten.Karteikasten
	)

	if err = decoder.Decode(&neu); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		panic(err)
	}

	if neu.Beschreibung == "" || neu.Kategorie == "" || neu.Titel == "" || neu.Unterkategorie == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if uuid, err = client.GetUUID(); err != nil {
		internalError(err, w)
	}

	kasten = karteikaesten.Karteikasten{
		ID:             uuid,
		Autor:          username,
		Name:           neu.Titel,
		Kategorie:      neu.Kategorie,
		Unterkategorie: neu.Unterkategorie,
		Beschreibung:   neu.Beschreibung,
		Public:         neu.Public,
		Karten:         []karteikaesten.Karteikarte{},
	}

	if err = kaesten.KastenAnlegen(kasten); err != nil {
		internalError(err, w)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, uuid)
}

// EditControllerBearbeiten controller 1 for edit Bearbeiten
func EditControllerBearbeiten(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	var (
		username = GetUser(w, r)
		data     Edit
		err      error
		kastenid = mux.Vars(r)["kastenid"]
	)

	kasten, err := karteikaesten.New().KastenByID(kastenid)
	if err != nil {
		internalError(err, w)
	}

	if kasten.Autor != username {
		forbidden(w)
		return
	}

	data = Edit{
		Titel:           kasten.Name,
		IstKategorie:    kasten.Kategorie,
		IstSubKategorie: kasten.Unterkategorie,
		Beschreibung:    kasten.Beschreibung,
		Public:          kasten.Public,
	}

	EditControllerBase(w, r, data)
}

// EditControllerNeu controller 1 for edit Neu
func EditControllerNeu(w http.ResponseWriter, r *http.Request) {
	EditControllerBase(w, r, Edit{})
}

// EditControllerBase base controller 1 for edit
func EditControllerBase(w http.ResponseWriter, r *http.Request, data Edit) {
	defer recoverInternalError()

	var err error

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	kats, err := kategorien.New().AlleKategorien()
	if err != nil {
		internalError(err, w)
	}

	for _, kat := range kats {
		kat := EditKategorie{
			Name: kat.ID,
			Sub:  kat.Unterkategorie,
		}

		data.Kategorien = append(data.Kategorien, kat)
	}

	customExecuteTemplate(w, r, templates.Edit, data)
}

// Edit2Controller controller 2 for edit
func Edit2Controller(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

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
			KastenID    string
		}
	)

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	var (
		err      error
		user     = GetUser(w, r)
		kastenid = mux.Vars(r)["kastenid"]
		data     Edit2
	)

	kasten, err := karteikaesten.New().KastenByID(kastenid)
	if err != nil {
		internalError(err, w)
	}

	data = Edit2{
		Name:      kasten.Name,
		Kategorie: kasten.Kategorie,
		SubKat:    kasten.Unterkategorie,
		Anzahl:    kasten.AnzahlKarten(),
		Karten:    []Karte{},
		KastenID:  kastenid,
	}

	if data.Fortschritt, err = kasten.Fortschritt(user); err != nil {
		internalError(err, w)
	}

	for nr, karte := range kasten.Karten {
		data.Karten = append(data.Karten, Karte{
			Nr:    nr + 1,
			Titel: karte.Titel,
		})
	}

	customExecuteTemplate(w, r, templates.Edit2, data)
}
