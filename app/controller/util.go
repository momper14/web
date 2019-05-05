package controller

import (
	"fmt"
	"html/template"
	"net/http"
	templates "web/template"

	"github.com/fatih/structs"
)

type Sidemenu struct {
	Kasten      int
	MeineKasten int
}

type Navbar struct {
	Name string
}

func customExecuteTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	var navbar, sidemenu string
	//navbar = templates.NavbarNoLogin
	//sidemenu = templates.SidemenuNoLogin
	navbar = templates.NavbarLogin
	sidemenu = templates.SidemenuLogin

	var dataTmp map[string]interface{}

	if structs.IsStruct(data) {
		dataTmp = structs.Map(data)
	} else {
		var ok bool
		dataTmp, ok = data.(map[string]interface{})
		if !ok {
			dataTmp = make(map[string]interface{})
		}
	}

	sm := Sidemenu{
		Kasten:      22,
		MeineKasten: 7,
	}

	nb := Navbar{
		Name: "Max Mustermann",
	}

	dataTmp["Sidemenu"] = sm
	dataTmp["Navbar"] = nb

	t, err := template.ParseFiles(templateName, navbar, sidemenu)
	if err != nil {
		fmt.Print(err)
	}
	t.Execute(w, dataTmp)
}
