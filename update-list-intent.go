package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//UpdateListNameIntent is an intent to update single list name
type UpdateListNameIntent struct {
	ListRepo ListRepository
}

//Enact function is for UpdateListNameIntent to update list through http
func (updateListIntent UpdateListNameIntent) Enact(w http.ResponseWriter, r *http.Request) {
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

	errors = updateListIntent.ListRepo.Update(list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusNoContent|http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
