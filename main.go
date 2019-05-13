package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Momper14/web/app/router"
	"github.com/urfave/negroni"
)

func main() {
	router := router.GetRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Println("Listening...")
	s := &http.Server{
		Addr:         ":3001",
		Handler:      n,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
