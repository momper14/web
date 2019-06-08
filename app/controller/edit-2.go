package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Momper14/web/app/model"
	"github.com/gorilla/mux"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client"
	"github.com/Momper14/weblib/client/karteikaesten"
)

type (
	// Edit2Karte Karte for Edit 2
	Edit2Karte struct {
		Nr    int
		Titel string
	}

	// Edit2Headline Headline for Edit 2
	Edit2Headline struct {
		Name        string
		Kategorie   string
		SubKat      string
		Fortschritt int
		Anzahl      int
	}

	// Edit2 Data for Edit 2
	Edit2 struct {
		Headline Edit2Headline
		Karten   []Edit2Karte
		KastenID string
		Titel    string
		Frage    string
		Antwort  string
	}
)

// Edit2ControllerPost controller 2 for edit POST
func Edit2ControllerPost(w http.ResponseWriter, r *http.Request) {
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
		neu      model.Karteikarte
		kasten   karteikaesten.Karteikasten
		karte    karteikaesten.Karteikarte
		kastenid = mux.Vars(r)["kastenid"]
	)

	if err = decoder.Decode(&neu); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		panic(err)
	}

	if neu.Titel == "" || neu.Frage == "" || neu.Antwort == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if kasten, err = kaesten.KastenByID(kastenid); err != nil {
		if _, ok := err.(client.NotFoundError); ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		internalError(err, w)
	}

	if kasten.Autor != username {
		forbidden(w)
		return
	}

	karte = karteikaesten.Karteikarte{
		Titel:   neu.Titel,
		Frage:   neu.Frage,
		Antwort: neu.Antwort,
	}

	if err = kasten.KarteHinzufuegen(karte); err != nil {
		internalError(err, w)
	}

	ok(w)
}

// Edit2ControllerPut controller 2 for edit PUT
func Edit2ControllerPut(w http.ResponseWriter, r *http.Request) {
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
		neu      model.Karteikarte
		kasten   karteikaesten.Karteikasten
		karte    karteikaesten.Karteikarte
		kastenid = mux.Vars(r)["kastenid"]
		index    int
	)

	if index, err = strconv.Atoi(mux.Vars(r)["karte"]); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if err = decoder.Decode(&neu); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		panic(err)
	}

	if neu.Titel == "" || neu.Frage == "" || neu.Antwort == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if kasten, err = kaesten.KastenByID(kastenid); err != nil {
		if _, ok := err.(client.NotFoundError); ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		internalError(err, w)
	}

	if index < 0 || index >= kasten.AnzahlKarten() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if kasten.Autor != username {
		forbidden(w)
		return
	}

	karte = karteikaesten.Karteikarte{
		Titel:   neu.Titel,
		Frage:   neu.Frage,
		Antwort: neu.Antwort,
	}

	if err = kasten.KarteAktualisieren(index, karte); err != nil {
		internalError(err, w)
	}

	ok(w)
}

// Edit2ControllerDelete controller 2 for edit delete
func Edit2ControllerDelete(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	var (
		username = GetUser(w, r)
		kaesten  = karteikaesten.New()
		err      error
		kasten   karteikaesten.Karteikasten
		kastenid = mux.Vars(r)["kastenid"]
		index    int
	)

	if index, err = strconv.Atoi(mux.Vars(r)["karte"]); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if kasten, err = kaesten.KastenByID(kastenid); err != nil {
		if _, ok := err.(client.NotFoundError); ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		internalError(err, w)
	}

	if index < 0 || index >= kasten.AnzahlKarten() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if kasten.Autor != username {
		forbidden(w)
		return
	}

	if err = kasten.KarteLoeschen(index); err != nil {
		internalError(err, w)
	}

	ok(w)
}

// Edit2Controller controller 2 for edit
func Edit2Controller(w http.ResponseWriter, r *http.Request) {
	Edit2ControllerBase(w, r, -1)
}

// Edit2ControllerMitKarte controller 2 vor edit with Karte
func Edit2ControllerMitKarte(w http.ResponseWriter, r *http.Request) {
	if index, err := strconv.Atoi(mux.Vars(r)["karte"]); err != nil || index == -1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		Edit2ControllerBase(w, r, index)
	}
}

// Edit2ControllerBase base controller 2 for edit
func Edit2ControllerBase(w http.ResponseWriter, r *http.Request, index int) {
	defer recoverInternalError()

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
		if _, ok := err.(client.NotFoundError); ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		internalError(err, w)
	}

	data = Edit2{
		Headline: Edit2Headline{
			Name:      kasten.Name,
			Kategorie: kasten.Kategorie,
			SubKat:    kasten.Unterkategorie,
			Anzahl:    kasten.AnzahlKarten()},
		Karten:   []Edit2Karte{},
		KastenID: kastenid,
	}

	if index < -1 || index >= kasten.AnzahlKarten() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if index >= 0 {
		k := kasten.Karten[index]
		data.Titel = k.Titel
		data.Frage = k.Frage
		data.Antwort = k.Antwort
	}

	if data.Headline.Fortschritt, err = kasten.Fortschritt(user); err != nil {
		internalError(err, w)
	}

	for nr, karte := range kasten.Karten {
		data.Karten = append(data.Karten, Edit2Karte{
			Nr:    nr + 1,
			Titel: karte.Titel,
		})
	}

	customExecuteTemplate(w, r, templates.Edit2, data)
}
