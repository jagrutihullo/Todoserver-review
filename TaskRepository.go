package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//task gorm model used by gorm
type TaskModel struct {
	gorm.Model
	Description string
	Status      string
	Deadline    time.Time
	LID         uint
}

//epository for task
type TaskRepository interface {
	Create(task Task, list List) []error

	Fetch(ID uint) (Task, []error)

	Update(task Task) []error

	Delete(ID uint) []error
}

//create task
func (taskModel TaskModel) Create(task Task, list List) []error {
	var errors []error
	var todo TodoList
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return []error{err}
	}

	errors = db.First(&todo, "id = ? and isnull(deleted_at)", list.ID).GetErrors()
	if len(errors) != 0 {
		return errors
	}

	taskModel.Description = task.Description
	taskModel.Status = task.Status
	taskModel.Deadline = task.Deadline
	taskModel.LID = list.ID
	errors = db.Create(&taskModel).GetErrors()
	return errors
}

//update task
func (taskModel TaskModel) Update(task Task) []error {
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return []error{err}
	}

	errors = db.Find(&taskModel, task.ID).GetErrors()
	if len(errors) != 0 {
		return errors
	}
	taskModel.Description = task.Description
	taskModel.Status = task.Status
	taskModel.Deadline = task.Deadline

	errors = db.Save(&taskModel).GetErrors()
	return errors
}

//delete task - gorm creates TIMESTAMP deleted_at, not actual delete
func (taskModel TaskModel) Delete(ID uint) []error {
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return []error{err}
	}

	errors = db.First(&taskModel, "id = ?", ID).GetErrors()
	if len(errors) != 0 {
		return errors
	}
	errors = db.Delete(&taskModel).GetErrors()
	return errors
}

//fetch given task by ID
func (taskModel TaskModel) Fetch(ID uint) (Task, []error) {
	var tempTask Task
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return tempTask, []error{err}
	}

	errors = db.First(&taskModel, "id = ? and isnull(deleted_at)", ID).Scan(&tempTask).GetErrors()
	return tempTask, errors
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
