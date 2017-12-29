package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//intent to update list name
type UpdateListNameIntent struct {
	ListRepo TodoListRepository
}

//update list name function
func (updateListIntent UpdateListNameIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var list List
	var errors error

	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	updateListIntent.ListRepo = GormListRepo{}
	errors = updateListIntent.ListRepo.Update(list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}
