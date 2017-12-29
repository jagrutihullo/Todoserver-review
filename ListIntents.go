package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	var errors []error

	fetchListIntent.ListRepo = TodoList{}
	params := mux.Vars(r)
	name := params["name"]
	list, errors = fetchListIntent.ListRepo.Fetch(name)

	w.Header().Set("Content-Type", "application/json")
	if len(errors) == 0 {
		listJSON, err := json.Marshal(list)

		if err != nil {
			json.NewEncoder(w).Encode(err)
			log.Fatal("Cannot encode to JSON -", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(listJSON)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		log.Fatal("List does not exist -", errors)
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
	var errors []error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Cannot read request body -", err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.Unmarshal(body, &list)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Cannot encode to JSON -", err)
	}
	createListIntent.ListRepo = TodoList{}
	errors = createListIntent.ListRepo.Create(list)
	if len(errors) != 0 {
		json.NewEncoder(w).Encode(errors)
		log.Fatal("Cannot create list -", errors)
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
	var errors []error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Cannot read request body -", err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.Unmarshal(body, &list)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Cannot encode to JSON ", err)
	}
	updateListIntent.ListRepo = TodoList{}
	errors = updateListIntent.ListRepo.Update(list)
	if len(errors) != 0 {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(errors)
		log.Fatal("Cannot update list -", errors)
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
	var errors []error

	params := mux.Vars(r)
	name := params["name"]

	deleteListIntent.ListRepo = TodoList{}
	errors = deleteListIntent.ListRepo.Delete(name)
	if len(errors) != 0 {
		json.NewEncoder(w).Encode(errors)
		log.Fatal("Cannot delete list -", errors)
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
	var errors []error

	fetchAllIntent.ListRepo = TodoList{}
	lists, errors = fetchAllIntent.ListRepo.FetchAll()

	w.Header().Set("Content-Type", "application/json")
	if len(errors) == 0 {
		listsJSON, err := json.Marshal(lists)

		if err != nil {
			json.NewEncoder(w).Encode(err)
			log.Fatal("Cannot encode to JSON -", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(listsJSON)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		log.Fatal("No List exist -", errors)
	}
	return
}
