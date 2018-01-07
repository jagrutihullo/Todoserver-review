package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//FetchListIntent is an intent to access list by name
type FetchListIntent struct {
	ListRepo ListRepository
}

//Enact function is for FetchListIntent to access list through http
func (fetchListIntent FetchListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		list   List
		errors error
	)

	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	list.ID = uint(i)
	list, errors = fetchListIntent.ListRepo.FetchByID(list)

	w.Header().Set("Content-Type", "application/json")
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusNoContent|http.StatusBadRequest)
	}
	listJSON, err := json.Marshal(list)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(listJSON)
}
