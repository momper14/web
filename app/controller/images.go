package controller

import (
	"net/http"
	"strings"

	"github.com/Momper14/web/app/url"
	"github.com/gorilla/mux"
)

// ImageController controller for Images
func ImageController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	if !IstEingeloggt(w, r) {
		forbidden(w)
		return
	}

	var (
		user  = GetUser(w, r)
		image = mux.Vars(r)["image"]
	)

	index := strings.LastIndex(image, ".")
	if user != image[:index] {
		forbidden(w)
		return
	}

	h := http.StripPrefix(url.ImagePath, http.FileServer(http.Dir("static/images")))
	h.ServeHTTP(w, r)
}
