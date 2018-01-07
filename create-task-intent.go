package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateTaskIntent is an intent to create task under list
type CreateTaskIntent struct {
	ListRepo ListRepository
}

//Enact function is for CreateTaskIntent to create task through http
func (createTask CreateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		list   List
		task   Task
		errors error
	)

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

	list.Tasks = make([]Task, 1)
	list.Tasks[0] = task
	errors = createTask.ListRepo.CreateTask(list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
