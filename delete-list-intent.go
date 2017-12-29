package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//intent to delete list & tasks under it
type DeleteListIntent struct {
	ListRepo TodoListRepository
}

//delete function
func (deleteListIntent DeleteListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var errors error

	params := mux.Vars(r)
	name := params["name"]

	w.Header().Set("Content-Type", "application/json")
	deleteListIntent.ListRepo = GormListRepo{}
	errors = deleteListIntent.ListRepo.Delete(name)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}
