package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//TodoListRepository is repository interface for list and task functions
type TodoListRepository interface {
	Create(todoEntity List) error

	Fetch(list List) (List, error)

	Update(todoEntity List) error

	Delete(list List) error

	FetchAll() ([]List, error)

	CreateTask(list List) error

	FetchTask(list List) (List, error)

	UpdateTask(list List) error

	DeleteTask(list List) error
}

//GormListRepo is a structure that implements TodoListRepo functions
type GormListRepo struct {
}

//Create is a Gorm function to create list
func (glr GormListRepo) Create(todoEntity List) error {
	var (
		error1    error
		errorsArr []error
		todoList  TodoList
	)

	db, err := gorm.Open("mysql", "root:root@/todoserver?"+
		"charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	if db.HasTable(&TodoList{}) == false {
		db.CreateTable(&TodoList{})
		db.CreateTable(&TaskModel{})
	}

	todoList.Name = todoEntity.Name
	errorsArr = db.Create(&todoList).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1

}

//Update is a Gorm function to update list name
func (glr GormListRepo) Update(todoEntity List) error {
	var (
		error1    error
		errorsArr []error
		todoList  TodoList
	)

	db, err := gorm.Open("mysql", "root:root@/todoserver?"+
		"charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	errorsArr = db.Find(&todoList, "id = ? and "+
		"isnull(deleted_at)", todoEntity.ID).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return error1
	}
	todoList.Name = todoEntity.Name
	errorsArr = db.Save(&todoList).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//Delete is a Gorm function to delete list and tasks under it
//Gorm creates deleted_at TIMESTAMP, it does not actually deletes record
func (glr GormListRepo) Delete(list List) error {
	var (
		tasks     []TaskModel
		error1    error
		errorsArr []error
		todoList  TodoList
	)

	db, err := gorm.Open("mysql", "root:root@/todoserver?"+
		"charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	errorsArr = db.First(&todoList, "id = ? and "+
		"isnull(deleted_at)", list.ID).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return error1
	}

	db.Find(&tasks, "l_id = ? and isnull(deleted_at)", todoList.ID)
	for i := range tasks {
		errorsArr = db.Delete(tasks[i]).GetErrors()
		if len(errorsArr) != 0 {
			error1 = ErrorsConv(errorsArr)
			return error1
		}
	}

	errorsArr = db.Delete(todoList).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//Fetch is a Gorm function to access list by name and tasks under it
func (glr GormListRepo) Fetch(list List) (List, error) {
	var (
		tempList  List
		error1    error
		errorsArr []error
		todoList  TodoList
		taskModel TaskModel
	)

	db, err := gorm.Open("mysql", "root:root@/todoserver?"+
		"charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return tempList, err
	}

	errorsArr = db.First(&todoList, "id = ? and "+
		"isnull(deleted_at) ", list.ID).Scan(&tempList).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return tempList, error1
	}

	errorsArr = db.Find(&taskModel, "l_id = ? and "+
		"isnull(deleted_at)", tempList.ID).Scan(&tempList.Tasks).GetErrors()
	//error1 = ErrorsConv(errorsArr)
	return tempList, error1
}

//FetchAll is a Gorm function to access all lists
func (glr GormListRepo) FetchAll() ([]List, error) {
	var (
		tempList  []List
		error1    error
		errorsArr []error
	)

	db, err := gorm.Open("mysql", "root:root@/todoserver?"+
		"charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return tempList, err
	}

	errorsArr = db.Table("todo_lists").Where("isnull(deleted_at)").Select("id, " +
		"name, created_at, updated_at").Scan(&tempList).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return tempList, error1

}
