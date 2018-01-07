package main

import (
	"fmt"
	"testing"
	"time"
)

var (
	list     List
	lists    []List
	gormRepo GormListRepo
	err      error
	blank    string
)

func TestFetchAll(t *testing.T) {

	fmt.Println("Hello")
	lists, err = gormRepo.FetchAll()
	if err.Error() != "" {
		t.Errorf("List Table empty, got %s, want list of lists", err)
	} else {
		for i := range lists {
			fmt.Println(lists[i])
		}
	}
}

func TestFetch(t *testing.T) {
	list = List{
		ID: 1,
	}
	list, err = gormRepo.FetchByID(list)
	fmt.Println(list)
	if err.Error() != "" {
		t.Errorf("List not fetched, got %s, want list with ID %d "+
			"and all tasks under it", err, list.ID)
	} else {
		fmt.Println(list)
	}
}

func TestCreate(t *testing.T) {
	list = List{
		Name: "demo list 1",
	}
	err = gormRepo.Create(list)
	if err.Error() != "" {
		t.Errorf("List not created, got %s, want no error", err)
	} else {
		fmt.Println("List Created")
	}
}

func TestUpdate(t *testing.T) {
	list = List{
		ID:   1,
		Name: "software engg list",
	}
	err = gormRepo.Update(list)
	if err.Error() != "" {
		t.Errorf("List not updated, got %s, want updated Name = %s", err, list.Name)
	} else {
		fmt.Println("Updated list is ")
		list, err = gormRepo.FetchByID(list)
		fmt.Println(list)
	}
}

func TestDelete(t *testing.T) {
	list = List{
		ID: 1,
	}
	err = gormRepo.Delete(list)
	if err.Error() != "" {
		t.Errorf("List not deleted, got %s, want no error", err)
	} else {
		list, err = gormRepo.FetchByID(list)
		fmt.Println("Correct Output after deletion: ", err)
	}
}

func TestCreateTask(t *testing.T) {

	list = List{
		ID: 1,
		Tasks: []Task{{
			Description: "task2 list1",
			Status:      "pending",
			Deadline:    time.Date(2018, 3, 15, 00, 00, 00, 00, time.UTC),
		},
		},
	}
	err = gormRepo.CreateTask(list)
	if err.Error() != "" {
		t.Errorf("Task not created, got %s, want no error", err)
	} else {
		fmt.Println("Task Created")
	}
}

func TestUpdateTask(t *testing.T) {
	list = List{
		Tasks: []Task{{
			ID:          1,
			Description: "todoserver testing",
			Status:      "pending",
			Deadline:    time.Date(2018, 3, 15, 00, 00, 00, 00, time.UTC),
		},
		},
	}
	err = gormRepo.UpdateTask(list)
	if err.Error() != "" {
		t.Errorf("Task not updated, got %s", err)
		t.Error("want updated task as ", list.Tasks[0])
	} else {
		fmt.Println("Updated task is ")
		list, err = gormRepo.FetchTaskByID(list)
		fmt.Println(list)
	}
}

func TestDeleteTask(t *testing.T) {
	list = List{
		Tasks: []Task{{
			ID: 1,
		},
		},
	}
	err = gormRepo.DeleteTask(list)
	if err.Error() != "" {
		t.Errorf("Task not deleted, got %s, want no error", err)
	} else {
		list, err = gormRepo.FetchTaskByID(list)
		fmt.Println("Correct Output after deletion: ", err)
	}
}

func TestFetchTask(t *testing.T) {
	list = List{
		Tasks: []Task{{
			ID: 1,
		},
		},
	}
	list, err = gormRepo.FetchTaskByID(list)
	if err != nil {
		t.Errorf("Task not fetched, got %s, want task with ID %d", err, list.Tasks[0].ID)
	} else {
		fmt.Println(list)
	}
}

func TestFetchN(t *testing.T) {
	list = List{
		ID: 0,
	}
	list, err = gormRepo.FetchByID(list)
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
		list, err = gormRepo.FetchByID(list)
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
		list, err = gormRepo.FetchByID(list)
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
		list, err = gormRepo.FetchTaskByID(list)
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
		list, err = gormRepo.FetchTaskByID(list)
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
	list, err = gormRepo.FetchTaskByID(list)
	if err.Error() != "" {
		fmt.Printf("Task not fetched, %s", err)
	} else {
		t.Errorf("Task fetched, got task with id- %d, want - record not found", list.Tasks[0].ID)
		fmt.Println(list)
	}
}
