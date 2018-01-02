package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//DeleteTaskIntent is an intent to delete task
type DeleteTaskIntent struct {
	ListRepo TodoListRepository
}

//Enact function is for DeleteTaskIntent to delete task through http
func (deleteTask DeleteTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		errors error
		list   List
	)

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	list.Tasks = make([]Task, 1)
	list.Tasks[0].ID = uint(i)
	errors = deleteTask.ListRepo.DeleteTask(list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
