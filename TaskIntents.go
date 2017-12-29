package main

import (
	"encoding/json"
	"io/ioutil"
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
	fetchTask.TaskRepo = TaskModel{}
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

	createTask.TaskRepo = TaskModel{}
	errors = createTask.TaskRepo.Create(task, list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

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

	updateTask.TaskRepo = TaskModel{}
	errors = updateTask.TaskRepo.Update(task)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

//intent to delete task
type DeleteTaskIntent struct {
	TaskRepo TaskRepository
}

//delete task function
func (deleteTask DeleteTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var errors error

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	deleteTask.TaskRepo = TaskModel{}
	errors = deleteTask.TaskRepo.Delete(uint(i))
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return

}
