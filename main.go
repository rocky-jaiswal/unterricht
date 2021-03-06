package main

import (
	"./controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Index)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
