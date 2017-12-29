package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//intent to update task
type UpdateTaskIntent struct {
	TaskRepo TaskRepository
}

//update task function
func (updateTask UpdateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var task Task
	var errors error

	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	updateTask.TaskRepo = GormTaskRepo{}
	errors = updateTask.TaskRepo.Update(task)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}
