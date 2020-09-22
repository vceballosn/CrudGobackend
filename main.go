package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID       int    `json:"ID"`
	NOMBRE   string `json:"NOMBRE"`
	APELLIDO string `json:"APELIIDO"`
	EMAIL    string `json:"EMAIL"`
}

type allTask []task

var tasks = allTask{
	{
		ID:       1,
		NOMBRE:   " VLADIMIR",
		APELLIDO: "CEBALLOS",
		EMAIL:    "VLADIMIR.CEBALLOS",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API! VLADIMIR CEBALLOS")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":3000", router))
}
