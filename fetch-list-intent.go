package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//intent to fetch list by name
type FetchListIntent struct {
	ListRepo TodoListRepository
}

//fetch list function
func (fetchListIntent FetchListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var list List
	var errors error

	fetchListIntent.ListRepo = GormListRepo{}
	params := mux.Vars(r)
	name := params["name"]
	list, errors = fetchListIntent.ListRepo.Fetch(name)

	w.Header().Set("Content-Type", "application/json")
	if errors == nil {
		listJSON, err := json.Marshal(list)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(listJSON)
	} else {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	}
	return
}
