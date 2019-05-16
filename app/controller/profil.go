package controller

import (
	"net/http"
	"time"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/Momper14/weblib/client/users"
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
