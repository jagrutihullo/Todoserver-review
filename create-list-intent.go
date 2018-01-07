package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CreateListIntent is an intent to create list
type CreateListIntent struct {
	ListRepo ListRepository
}

//Enact function is for CreateListIntent to create list through http
func (createListIntent CreateListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		list   List
		errors error
	)

	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	errors = createListIntent.ListRepo.Create(list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
