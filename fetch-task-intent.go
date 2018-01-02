package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//FetchTaskIntent is an intent to access task
type FetchTaskIntent struct {
	ListRepo TodoListRepository
}

//Enact function is for FetchTaskIntent to fetch task through http
func (fetchTask FetchTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		list   List
		errors error
	)

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	list.Tasks = make([]Task, 1)
	list.Tasks[0].ID = uint(i)
	list, errors = fetchTask.ListRepo.FetchTask(list)

	if errors != nil {
		http.Error(w, errors.Error(), http.StatusNoContent|http.StatusBadRequest)
	}
	taskJSON, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(taskJSON)
}
