package main

import (
	"log"

	"web/app/router"

	"github.com/urfave/negroni"
)

func main() {
	router := router.GetRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Println("Listening:")
	n.Run(":3001")
}
