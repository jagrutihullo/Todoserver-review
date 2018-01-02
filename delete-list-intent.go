package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//DeleteListIntent is an intent to delete list & tasks under it
type DeleteListIntent struct {
	ListRepo TodoListRepository
}

//Enact function is for DeleteListIntent to delete list through http
func (deleteListIntent DeleteListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		errors error
		list   List
	)

	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	list.ID = uint(i)
	errors = deleteListIntent.ListRepo.Delete(list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
