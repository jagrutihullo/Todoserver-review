package main

import (
	"fmt"
	"testing"
	"time"
)

func TestFetchN(t *testing.T) {
	list = List{
		ID: 0,
	}
	list, err = gormRepo.Fetch(list)
	if err.Error() != "" {
		fmt.Printf("List not fetched, %s", err)
	} else {
		t.Errorf("List fetched, got list with id- %d, want - record not found", list.ID)
		fmt.Println(list)
	}
}

func TestCreateN(t *testing.T) {
	list = List{
		Name: "demo list",
	}
	err = gormRepo.Create(list)
	if err.Error() != "" {
		fmt.Printf("List not created, %s", err)
	} else {
		t.Errorf("List Created, got no error, want - duplicate key error")
	}
}

func TestUpdateN(t *testing.T) {
	list = List{
		ID:   0,
		Name: "demo list 2",
	}
	err = gormRepo.Update(list)
	if err.Error() != "" {
		fmt.Printf("List not updated, %s", err)
	} else {
		t.Errorf("got Updated list, want - record not found")
		list, err = gormRepo.Fetch(list)
		fmt.Println(list)
	}
}

func TestDeleteN(t *testing.T) {
	list = List{
		ID: 2,
	}
	err = gormRepo.Delete(list)
	if err.Error() != "" {
		fmt.Printf("List not deleted, %s", err)
	} else {
		list, err = gormRepo.Fetch(list)
		t.Errorf("List deleted, got %s, want list with ID - %d", err, list.ID)
	}
}

func TestCreateTaskN(t *testing.T) {
	list = List{
		ID: 6,
		Tasks: []Task{{
			Description: "task1 list6",
			Status:      "pending",
			Deadline:    time.Date(2018, 3, 15, 00, 00, 00, 00, time.UTC),
		},
		},
	}
	err = gormRepo.CreateTask(list)
	if err.Error() != "" {
		fmt.Printf("Task not created, %s", err)
	} else {
		t.Errorf("Task Created, got no error, want - record not found for list id - %d", list.ID)
	}
}

func TestUpdateTaskN(t *testing.T) {
	list = List{
		Tasks: []Task{{
			ID:          0,
			Description: "todoserver testing",
			Status:      "pending",
			Deadline:    time.Date(2018, 3, 15, 00, 00, 00, 00, time.UTC),
		},
		},
	}
	err = gormRepo.UpdateTask(list)
	if err.Error() != "" {
		fmt.Printf("Task not updated, %s", err)
	} else {
		t.Errorf("got Updated task, want - record not found")
		list, err = gormRepo.FetchTask(list)
		fmt.Println(list)
	}
}

func TestDeleteTaskN(t *testing.T) {
	list = List{
		Tasks: []Task{{
			ID: 1,
		},
		},
	}
	err = gormRepo.DeleteTask(list)
	if err.Error() != "" {
		fmt.Printf("Task not deleted, %s", err)
	} else {
		list, err = gormRepo.FetchTask(list)
		t.Errorf("Task deleted, got %s, want task with ID - %d", err, list.Tasks[0].ID)
	}
}

func TestFetchTaskN(t *testing.T) {
	list = List{
		Tasks: []Task{{
			ID: 0,
		},
		},
	}
	list, err = gormRepo.FetchTask(list)
	if err.Error() != "" {
		fmt.Printf("Task not fetched, %s", err)
	} else {
		t.Errorf("Task fetched, got task with id- %d, want - record not found", list.Tasks[0].ID)
		fmt.Println(list)
	}
}
