package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
	deleteTask.TaskRepo = GormTaskRepo{}
	errors = deleteTask.TaskRepo.Delete(uint(i))
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return

}
