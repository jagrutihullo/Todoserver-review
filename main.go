package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//function to convert array of errors into single error
func ErrorsConv(errorsArr []error) error {
	var error1 error
	var errString string

	for i := range errorsArr {
		errString += errorsArr[i].Error()
	}
	error1 = errors.New(errString)
	return error1
}

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
