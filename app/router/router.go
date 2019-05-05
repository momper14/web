package router

import (
	"net/http"
	"web/app/controller"
	"web/app/url"

	"github.com/gorilla/mux"
)

// GetRouter returns all routers
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix(url.StaticPath).Handler(
		http.StripPrefix(url.StaticPath, http.FileServer(http.Dir("static"))))

	router.HandleFunc(url.ViewPath, controller.ViewController).Methods("GET")

	router.HandleFunc(url.ProfilPath, controller.ProfilController).Methods("GET")

	router.HandleFunc(url.RegisterPath, controller.RegisterController).Methods("GET")

	router.HandleFunc(url.MeineKarteienPath, controller.MeineKarteienController).Methods("GET")

	router.HandleFunc(url.LernPath, controller.LernController).Methods("GET")

	router.HandleFunc(url.KarteikastenPath, controller.KarteikastenController).Methods("GET")

	router.HandleFunc(url.Edit2Path, controller.Edit2Controller).Methods("GET")
	router.HandleFunc(url.EditPath, controller.EditController).Methods("GET")

	router.HandleFunc(url.HomePath, controller.IndexController).Methods("GET")
	return router
}
