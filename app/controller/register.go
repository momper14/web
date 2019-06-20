package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Momper14/web/app/model"
	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/users"
	userspkg "github.com/Momper14/weblib/client/users"
	"github.com/gorilla/mux"
)

// RegisterController controller for register
func RegisterController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()
	data := make(map[string]interface{})

	customExecuteTemplate(w, r, templates.Register, data)
}

// RegisterControllerPost controller for register Post
func RegisterControllerPost(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	var (
		vorhanden     bool
		err           error
		registrierung model.Registrierung
		decoder       = json.NewDecoder(r.Body)
	)

	if err = decoder.Decode(&registrierung); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		panic(err)
	}

	if l := len(registrierung.Passwort); l != 128 {
		http.Error(w, fmt.Sprintf("Password: %s hat eine ungültige SHA-512 Zeichenlänge von %d", registrierung.Passwort, l), http.StatusBadRequest)
		return
	}

	if !registrierung.Akzeptiert {
		http.Error(w, "Datenschutzerklärung muss Akzeptiert werden", http.StatusForbidden)
		return
	}

	users := userspkg.New()

	if vorhanden, err = users.CheckName(registrierung.Name); err != nil {
		internalError(err, w)
	} else if vorhanden {
		http.Error(w, "Username bereits vergeben", http.StatusConflict)
		return
	}

	if vorhanden, err = users.CheckEmail(registrierung.EMail); err != nil {
		internalError(err, w)
	} else if vorhanden {
		http.Error(w, "Email bereits vergeben", http.StatusConflict)
		return
	}

	fileName := fmt.Sprintf("%s/%s.svg", imagePath, registrierung.Name)

	users.UserErstellen(userspkg.User{
		Name:     registrierung.Name,
		Password: registrierung.Passwort,
		Email:    registrierung.EMail,
		Bild:     fileName,
		Seit:     time.Now().Unix(),
	})

	file, err := ioutil.ReadFile("static/res/Icons/Mein-Profil.svg")
	errF(err, w)
	errF(ioutil.WriteFile(fileName[1:], file, 0644), w)

	session, err := store.Get(r, SessionCookieName)
	errF(err, w)

	session.Values["authenticated"] = true
	session.Values["user"] = registrierung.Name
	session.Save(r, w)

	ok(w)
}

// RegisterControllerCheckName controller for checking Name
func RegisterControllerCheckName(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	name := mux.Vars(r)["name"]
	users := users.New()

	vorhanden, err := users.CheckName(name)
	errF(err, w)

	if vorhanden {
		http.Error(w, "Nutzername vergeben", http.StatusConflict)
		return
	}

	ok(w)
}

// RegisterControllerCheckEMail controller for checking Name
func RegisterControllerCheckEMail(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	email := mux.Vars(r)["email"]
	users := users.New()

	vorhanden, err := users.CheckEmail(email)
	errF(err, w)

	if vorhanden {
		http.Error(w, "Email vergeben", http.StatusConflict)
		return
	}

	ok(w)
}
