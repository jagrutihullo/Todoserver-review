package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	ListRepo := GormListRepo{}

	fetchList := FetchListIntent{ListRepo}
	createList := CreateListIntent{ListRepo}
	updateList := UpdateListNameIntent{ListRepo}
	deleteList := DeleteListIntent{ListRepo}
	fetchAll := FetchAllListIntent{ListRepo}

	fetchTask := FetchTaskIntent{ListRepo}
	createTask := CreateTaskIntent{ListRepo}
	updateTask := UpdateTaskIntent{ListRepo}
	deleteTask := DeleteTaskIntent{ListRepo}

	router := mux.NewRouter()

	router.HandleFunc("/list/{id}", fetchList.Enact).Methods("GET")
	router.HandleFunc("/list", createList.Enact).Methods("POST")
	router.HandleFunc("/list", updateList.Enact).Methods("PATCH")
	router.HandleFunc("/list/{id}", deleteList.Enact).Methods("DELETE")
	router.HandleFunc("/list", fetchAll.Enact).Methods("GET")

	router.HandleFunc("/task/{id}", fetchTask.Enact).Methods("GET")
	router.HandleFunc("/list/{id}/task", createTask.Enact).Methods("POST")
	router.HandleFunc("/task", updateTask.Enact).Methods("PATCH")
	router.HandleFunc("/task/{id}", deleteTask.Enact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
	http.Handle("/", router)
}
