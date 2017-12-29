package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	var errors []error

	fetchTask.TaskRepo = TaskModel{}
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Cannot identify task id -", err)
	}

	task, errors = fetchTask.TaskRepo.Fetch(uint(i))

	w.Header().Set("Content-Type", "application/json")
	if len(errors) == 0 {
		taskJSON, err := json.Marshal(task)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			log.Fatal("Cannot encode to JSON -", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(taskJSON)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		log.Fatal("Task does not exist -", errors)
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
	var errors []error

	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	fmt.Println(err)
	list.ID = uint(i)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Cannot read request body -", err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.Unmarshal(body, &task)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Cannot encode to JSON -", err)
	}

	createTask.TaskRepo = TaskModel{}
	errors = createTask.TaskRepo.Create(task, list)
	if len(errors) != 0 {
		json.NewEncoder(w).Encode(errors)
		log.Fatal("Cannot create task under list -", errors)
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
	var errors []error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Cannot read request body -", err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.Unmarshal(body, &task)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Cannot encode to JSON -", err)
	}

	updateTask.TaskRepo = TaskModel{}
	errors = updateTask.TaskRepo.Update(task)
	if len(errors) != 0 {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(errors)
		log.Fatal("Cannot update task -", errors)
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
	var errors []error

	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	fmt.Println(err)
	deleteTask.TaskRepo = TaskModel{}
	errors = deleteTask.TaskRepo.Delete(uint(i))
	if len(errors) != 0 {
		json.NewEncoder(w).Encode(errors)
		log.Fatal("Cannot delete task -", errors)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return

}
