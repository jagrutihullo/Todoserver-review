package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//UpdateTaskIntent is an intent to update single task
type UpdateTaskIntent struct {
	ListRepo ListRepository
}

//Enact function is for UpdateTaskIntent to update task through http
func (updateTask UpdateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		task   Task
		errors error
		list   List
	)

	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	list.Tasks = make([]Task, 1)
	list.Tasks[0] = task
	errors = updateTask.ListRepo.UpdateTask(list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusNoContent|http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
