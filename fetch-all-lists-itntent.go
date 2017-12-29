package main

import (
	"encoding/json"
	"net/http"
)

//intent to fetch all lists
type FetchAllListIntent struct {
	ListRepo TodoListRepository
}

//fetch lists function
func (fetchAllIntent FetchAllListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var lists []List
	var errors error

	fetchAllIntent.ListRepo = GormListRepo{}
	lists, errors = fetchAllIntent.ListRepo.FetchAll()

	w.Header().Set("Content-Type", "application/json")
	if errors == nil {
		listsJSON, err := json.Marshal(lists)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(listsJSON)
	} else {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	}
	return
}
