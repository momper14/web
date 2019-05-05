package controller

import (
	"net/http"
	templates "web/template"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	customExecuteTemplate(w, r, templates.Register, data)
}
