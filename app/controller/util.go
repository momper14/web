package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Momper14/web/templates"
	"github.com/Momper14/weblib/client/karteikaesten"
	"github.com/fatih/structs"
)

func customExecuteTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	var navbar, sidemenu string

	type Sidemenu struct {
		Kasten      int
		MeineKasten int
	}

	type Navbar struct {
		Name string
	}

	//navbar = templates.NavbarNoLogin
	//sidemenu = templates.SidemenuNoLogin
	navbar = templates.NavbarLogin
	sidemenu = templates.SidemenuLogin

	var (
		dataTmp map[string]interface{}
		sm      Sidemenu
		nb      Navbar
		err     error
	)

	if structs.IsStruct(data) {
		dataTmp = structs.Map(data)
	} else {
		var ok bool
		dataTmp, ok = data.(map[string]interface{})
		if !ok {
			dataTmp = make(map[string]interface{})
		}
	}

	if sm.Kasten, err = karteikaesten.New().AnzahlOeffentlicherKaesten(); err != nil {
		internalError(err, w, r)
	}

	if sm.MeineKasten, err = karteikaesten.New().AnzahlKaestenUser(GetUser()); err != nil {
		internalError(err, w, r)
	}

	nb.Name = GetUser()

	dataTmp["Sidemenu"] = sm
	dataTmp["Navbar"] = nb

	t, err := template.ParseFiles(templateName, navbar, sidemenu)
	if err != nil {
		fmt.Print(err)
	}
	t.Execute(w, dataTmp)
}

// GetUser Returns the current User
func GetUser() string {
	return "Max Mustermann"
}

func internalError(err error, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))

	panic(err)
}

func recoverInternalError() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
