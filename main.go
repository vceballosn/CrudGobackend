package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		EMAIL:    "VLADIMIR.CEBALLOS@gmail.com",
	},
}

func getTaks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Insert a valid data")
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API! VLADIMIR CEBALLOS THE BEST DEVELOPER")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tasks", getTaks).Methods("GET")
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", createTask).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
