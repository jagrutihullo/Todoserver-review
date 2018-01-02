package main

import (
	"encoding/json"
	"net/http"
)

//FetchAllListIntent is an intent to access all lists
type FetchAllListIntent struct {
	ListRepo TodoListRepository
}

//Enact function is for FetchAllListIntent to access lists through http
func (fetchAllIntent FetchAllListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		lists  []List
		errors error
	)

	lists, errors = fetchAllIntent.ListRepo.FetchAll()

	w.Header().Set("Content-Type", "application/json")
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusNoContent|http.StatusBadRequest)
	}
	listsJSON, err := json.Marshal(lists)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(listsJSON)
}
