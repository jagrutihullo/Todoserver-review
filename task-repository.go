package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//repository for task
type TaskRepository interface {
	Create(task Task, list List) error

	Fetch(ID uint) (Task, error)

	Update(task Task) error

	Delete(ID uint) error
}

type GormTaskRepo struct {
}

//create task
func (gtr GormTaskRepo) Create(task Task, list List) error {
	var error1 error
	var errorsArr []error
	var todo TodoList
	var taskModel TaskModel

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	errorsArr = db.First(&todo, "id = ? and isnull(deleted_at)", list.ID).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return error1
	}
	taskModel = TaskToModel(task)
	taskModel.LID = list.ID
	errorsArr = db.Create(&taskModel).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//update task
func (gtr GormTaskRepo) Update(task Task) error {
	var error1 error
	var errorsArr []error
	var taskModel TaskModel

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	errorsArr = db.First(&taskModel, task.ID).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return error1
	}
	taskModel = TaskToModel(task)
	errorsArr = db.Save(&taskModel).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//delete task - gorm creates TIMESTAMP deleted_at, not actual delete
func (gtr GormTaskRepo) Delete(ID uint) error {
	var errorsArr []error
	var error1 error
	var taskModel TaskModel

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	errorsArr = db.First(&taskModel, "id = ?", ID).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return error1
	}
	errorsArr = db.Delete(&taskModel).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//fetch given task by ID
func (gtr GormTaskRepo) Fetch(ID uint) (Task, error) {
	var tempTask Task
	var error1 error
	var errorsArr []error
	var taskModel TaskModel

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return tempTask, err
	}

	errorsArr = db.First(&taskModel, "id = ? and isnull(deleted_at)", ID).Scan(&tempTask).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return tempTask, error1
}

// func main() {
// 	var taske TaskEntity
// 	var task Task
// 	var errors []error

// 	taske, errors = task.Fetch(1)
// 	fmt.Println(taske)
// 	fmt.Println(errors)

// 	taske = TaskEntity{
// 		ID:          1,
// 		Description: "task1 list1",
// 		Status:      "pending",
// 		Deadline:    time.Date(2018, 3, 15, 00, 00, 00, 00, time.UTC),
// 		LID:         1,
// 	}
// 	errors = task.Update(taske)
// 	fmt.Println(errors)

// 	errors = task.Delete(4)
// 	fmt.Println(errors)
// }
