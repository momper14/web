package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Momper14/web/app/model"
	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/Momper14/weblib/client/users"
	"github.com/gorilla/mux"
)

// ProfilController controller for Profile
func ProfilController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	const layout = "02.01.2006"

	type Data struct {
		Bild     string
		Name     string
		Email    string
		Karten   int
		Karteien int
		Seit     string
	}

	var (
		err    error
		data   Data
		userid string
	)

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	userid = GetUser(w, r)

	user, err := users.New().UserByID(userid)
	if err != nil {
		internalError(err, w)
	}

	data = Data{
		Bild:  user.Bild,
		Name:  user.Name,
		Email: user.Email,
		Seit:  time.Unix(user.Seit, 0).Format(layout),
	}

	if data.Karteien, err = karteikaesten.New().AnzahlKaestenUser(userid); err != nil {
		internalError(err, w)
	}

	if data.Karten, err = karteikaesten.New().AnzahlKartenUser(userid); err != nil {
		internalError(err, w)
	}

	customExecuteTemplate(w, r, templates.Profil, data)
}

// ProfilControllerPut controller for profile Put
func ProfilControllerPut(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	var (
		name    = GetUser(w, r)
		users   = users.New()
		err     error
		decoder = json.NewDecoder(r.Body)
		update  model.UpdateProfil
	)

	if err := decoder.Decode(&update); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		panic(err)
	}

	if update.EMail == "" && update.Neu == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := users.UserByID(name)
	if err != nil {
		internalError(err, w)
	}

	if update.EMail != "" {

		vorhanden, err := users.CheckEmail(update.EMail)
		if err != nil {
			internalError(err, w)
		}

		if vorhanden {
			http.Error(w, "Email vergeben", http.StatusConflict)
			return
		}

		user.Email = update.EMail
	}

	if update.Neu != "" {

		if l := len(update.Neu); l != 128 {
			http.Error(w, fmt.Sprintf("Password: %s hat eine ungültige SHA-512 Zeichenlänge von %d", update.Neu, l), http.StatusBadRequest)
			return
		}

		if update.Passwort != user.Password {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if update.Passwort == update.Neu {
			http.Error(w, "Neues Passwort identisch dem alten", http.StatusConflict)
			return
		}

		user.Password = update.Neu
	}

	if err := users.UserAktualisieren(user); err != nil {
		internalError(err, w)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

// ProfilControllerDelete controller for profile delete
func ProfilControllerDelete(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	name := GetUser(w, r)
	users := users.New()

	if err := users.UserLoeschen(name); err != nil {
		internalError(err, w)
	}

	session, err := store.Get(r, SessionCookieName)
	if err != nil {
		internalError(err, w)
	}

	delete(session.Values, "authenticated")
	delete(session.Values, "user")
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Profil gelöscht")
}

// ProfilControllerCheckPasswort controller for checking Name
func ProfilControllerCheckPasswort(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	passwort := mux.Vars(r)["passwort"]
	name := GetUser(w, r)
	users := users.New()
	user, err := users.UserByID(name)
	if err != nil {
		internalError(err, w)
	}

	if l := len(passwort); l != 128 {
		http.Error(w, fmt.Sprintf("Password: %s hat eine ungültige SHA-512 Zeichenlänge von %d", passwort, l), http.StatusBadRequest)
		return
	}

	if passwort == user.Password {
		http.Error(w, "Passwort darf nicht dem alten entsprechen", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

// ProfilControllerCheckEMail controller for checking Name
func ProfilControllerCheckEMail(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	email := mux.Vars(r)["email"]
	users := users.New()

	vorhanden, err := users.CheckEmail(email)
	if err != nil {
		internalError(err, w)
	}

	if vorhanden {
		http.Error(w, "Email vergeben", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}
