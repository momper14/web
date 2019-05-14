package controller

import (
	"net/http"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/Momper14/weblib/client/users"
)

// IndexController controller for index-page
func IndexController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	type Index struct {
		Nutzer, Karten, Karteien int
	}

	var (
		data Index
		err  error
	)

	if data.Karteien, err = karteikaesten.New().AnzahlOeffentlicherKaesten(); err != nil {
		internalError(err, w)
	}
	if data.Karten, err = karteikaesten.New().AnzahlOeffentlicherKarten(); err != nil {
		internalError(err, w)
	}
	if data.Nutzer, err = users.New().AnzahlUsers(); err != nil {
		internalError(err, w)
	}

	customExecuteTemplate(w, r, templates.Index, data)
}
