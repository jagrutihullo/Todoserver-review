package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//CreateTask is a Gorm function to create task under a list
func (glr GormListRepo) CreateTask(list List) error {
	var (
		error1    error
		errorsArr []error
		//todo      TodoList
		taskModel TaskModel
		tempList  List
	)

	db, err := gorm.Open("mysql", "root:root@/todoserver?"+
		"charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	tempList, error1 = glr.Fetch(list)
	if error1 != nil {
		return error1
	}
	taskModel = TaskToModel(list.Tasks[0], taskModel)
	taskModel.LID = tempList.ID
	errorsArr = db.Create(&taskModel).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//UpdateTask is a Gorm function to update task
func (glr GormListRepo) UpdateTask(list List) error {
	var (
		error1    error
		errorsArr []error
		taskModel TaskModel
	)

	db, err := gorm.Open("mysql", "root:root@/todoserver?"+
		"charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	errorsArr = db.First(&taskModel, "id = ? and "+
		"isnull(deleted_at)", list.Tasks[0].ID).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return error1
	}
	taskModel = TaskToModel(list.Tasks[0], taskModel)
	errorsArr = db.Save(&taskModel).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//DeleteTask is a Gorm function to delete task
//Gorm creates deleted_at TIMESTAMP, it does not actually deletes record
func (glr GormListRepo) DeleteTask(list List) error {
	var (
		errorsArr []error
		error1    error
		taskModel TaskModel
	)

	db, err := gorm.Open("mysql", "root:root@/todoserver?"+
		"charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	errorsArr = db.First(&taskModel, "id = ? and "+
		"isnull(deleted_at)", list.Tasks[0].ID).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return error1
	}
	errorsArr = db.Delete(&taskModel).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//FetchTask is a Gorm function to access task by id
func (glr GormListRepo) FetchTask(list List) (List, error) {
	var (
		tempList  List
		error1    error
		errorsArr []error
		taskModel TaskModel
	)

	db, err := gorm.Open("mysql", "root:root@/todoserver?"+
		"charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return tempList, err
	}

	tempList.Tasks = make([]Task, 1)

	errorsArr = db.First(&taskModel, "id = ? and "+
		"isnull(deleted_at)", list.Tasks[0].ID).Scan(&tempList.Tasks[0]).GetErrors()
	error1 = ErrorsConv(errorsArr)
	tempList.ID = taskModel.LID
	list, error1 = glr.Fetch(tempList)
	if error1 != nil {
		return list, error1
	}

	tempList.Name = list.Name
	tempList.CreatedAt = list.CreatedAt
	tempList.UpdatedAt = list.UpdatedAt

	return tempList, error1
}
