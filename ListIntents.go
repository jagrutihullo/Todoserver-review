package main

import (
	"encoding/json"
	"io/ioutil"
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

	fetchListIntent.ListRepo = TodoList{}
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

//intent to create list
type CreateListIntent struct {
	ListRepo TodoListRepository
}

//create list function
func (createListIntent CreateListIntent) Enact(w http.ResponseWriter, r *http.Request) {
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
	createListIntent.ListRepo = TodoList{}
	errors = createListIntent.ListRepo.Create(list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

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
	updateListIntent.ListRepo = TodoList{}
	errors = updateListIntent.ListRepo.Update(list)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

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
	deleteListIntent.ListRepo = TodoList{}
	errors = deleteListIntent.ListRepo.Delete(name)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

//intent to fetch all lists
type FetchAllListIntent struct {
	ListRepo TodoListRepository
}

//fetch lists function
func (fetchAllIntent FetchAllListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var lists []List
	var errors error

	fetchAllIntent.ListRepo = TodoList{}
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
