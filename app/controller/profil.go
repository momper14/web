package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Momper14/web/app/model"
	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/Momper14/weblib/client/users"
	"github.com/gorilla/mux"
	"github.com/vincent-petithory/dataurl"
)

const imagePath = "/static/images"

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
	errF(err, w)

	data = Data{
		Name:  user.Name,
		Email: user.Email,
		Seit:  time.Unix(user.Seit, 0).Format(layout),
	}

	data.Bild, err = imageLastmod(user.Bild)
	errF(err, w)

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
		dataURL *dataurl.DataURL
		file    string
	)

	if err = decoder.Decode(&update); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		panic(err)
	}

	if update.EMail == "" && update.Neu == "" && update.Bild == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := users.UserByID(name)
	errF(err, w)

	if update.EMail != "" {

		vorhanden, err := users.CheckEmail(update.EMail)
		errF(err, w)

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

	if update.Bild != "" {
		if dataURL, err = dataurl.DecodeString(update.Bild); err != nil {
			internalError(err, w)
		}

		switch dataURL.ContentType() {
		case "image/png":
			file = fmt.Sprintf("%s/%s.png", imagePath, name)
			break
		case "image/jpeg":
			file = fmt.Sprintf("%s/%s.jpg", imagePath, name)
			break
		case "image/svg+xml":
			file = fmt.Sprintf("%s/%s.svg", imagePath, name)
			break
		case "image/vnd.microsoft.icon":
			file = fmt.Sprintf("%s/%s.ico", imagePath, name)
			break
		default:
			http.Error(w, "ungültiges Bildformat", http.StatusBadRequest)
			return
		}
		tmp := user.Bild
		errF(ioutil.WriteFile(file[1:], dataURL.Data, 0644), w)
		if tmp != file {
			errF(os.Remove(user.Bild[1:]), w)
		}
		user.Bild = file
	}

	errF(users.UserAktualisieren(user), w)

	ok(w)
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

	user, err := users.UserByID(name)
	errF(err, w)

	errF(users.UserLoeschen(name), w)

	session, err := store.Get(r, SessionCookieName)
	errF(err, w)

	delete(session.Values, "authenticated")
	delete(session.Values, "user")
	session.Save(r, w)

	errF(os.Remove(user.Bild[1:]), w)

	ok(w)
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
	errF(err, w)

	if l := len(passwort); l != 128 {
		http.Error(w, fmt.Sprintf("Password: %s hat eine ungültige SHA-512 Zeichenlänge von %d", passwort, l), http.StatusBadRequest)
		return
	}

	if passwort == user.Password {
		http.Error(w, "Passwort darf nicht dem alten entsprechen", http.StatusConflict)
		return
	}

	ok(w)
}

// ProfilControllerCheckEMail controller for checking Name
func ProfilControllerCheckEMail(w http.ResponseWriter, r *http.Request) {
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
