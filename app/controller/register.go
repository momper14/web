package controller

import (
	"net/http"

	"github.com/Momper14/web/templates"
)

// RegisterController controller for register
func RegisterController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()
	data := make(map[string]interface{})

	customExecuteTemplate(w, r, templates.Register, data)
}
