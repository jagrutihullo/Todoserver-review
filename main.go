package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var fetchList FetchListIntent
	var createList CreateListIntent
	var updateList UpdateListNameIntent
	var deleteList DeleteListIntent
	var fetchAll FetchAllListIntent

	var fetchTask FetchTaskIntent
	var createTask CreateTaskIntent
	var updateTask UpdateTaskIntent
	var deleteTask DeleteTaskIntent

	router := mux.NewRouter()

	router.HandleFunc("/list/{name}", fetchList.Enact).Methods("GET")
	router.HandleFunc("/list", createList.Enact).Methods("POST")
	router.HandleFunc("/list", updateList.Enact).Methods("PATCH")
	router.HandleFunc("/list/{name}", deleteList.Enact).Methods("DELETE")
	router.HandleFunc("/list", fetchAll.Enact).Methods("GET")

	router.HandleFunc("/task/{id}", fetchTask.Enact).Methods("GET")
	router.HandleFunc("/list/{id}/task", createTask.Enact).Methods("POST")
	router.HandleFunc("/task", updateTask.Enact).Methods("PATCH")
	router.HandleFunc("/task/{id}", deleteTask.Enact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
	http.Handle("/", router)
}
