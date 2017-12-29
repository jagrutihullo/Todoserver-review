package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//intent to fetch task
type FetchTaskIntent struct {
	TaskRepo TaskRepository
}

//fetch task function
func (fetchTask FetchTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var task Task
	var errors error

	w.Header().Set("Content-Type", "application/json")
	fetchTask.TaskRepo = GormTaskRepo{}
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	task, errors = fetchTask.TaskRepo.Fetch(uint(i))

	if errors == nil {
		taskJSON, err := json.Marshal(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(taskJSON)
	} else {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	}
	return
}
