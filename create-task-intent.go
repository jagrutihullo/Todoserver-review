package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//intent to create task under list
type CreateTaskIntent struct {
	TaskRepo TaskRepository
}

//create task function
func (createTask CreateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var list List
	var task Task
	var errors error

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	list.ID = uint(i)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if task.Description == "" {
		http.Error(w, "Task description cannot be empty", http.StatusBadRequest)
		return
	}

	createTask.TaskRepo = GormTaskRepo{}
	errors = createTask.TaskRepo.Create(task, list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}
