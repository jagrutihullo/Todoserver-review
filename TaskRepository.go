package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//task gorm model used by gorm
type Task struct {
	gorm.Model
	Description string
	Status      string
	Deadline    time.Time
	LID         uint
}

//epository for task
type TaskRepository interface {
	Create(taskEntity TaskEntity) []error

	Fetch(ID uint) (TaskEntity, []error)

	Update(taskEntity TaskEntity) []error

	Delete(ID uint) []error
}

//create task
func (task Task) Create(taskEntity TaskEntity) []error {
	var errors []error
	var todo TodoList
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return []error{err}
	}

	errors = db.First(&todo, "id = ? and isnull(deleted_at)", taskEntity.LID).GetErrors()
	if len(errors) != 0 {
		return errors
	}

	task.Description = taskEntity.Description
	task.Status = taskEntity.Status
	task.Deadline = taskEntity.Deadline
	task.LID = taskEntity.LID
	errors = db.Create(&task).GetErrors()
	return errors
}

//update task
func (task Task) Update(taskEntity TaskEntity) []error {
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return []error{err}
	}

	errors = db.Find(&task, taskEntity.ID).GetErrors()
	if len(errors) != 0 {
		return errors
	}
	task.Description = taskEntity.Description
	task.Status = taskEntity.Status
	task.Deadline = taskEntity.Deadline
	task.LID = taskEntity.LID

	errors = db.Save(&task).GetErrors()
	return errors
}

//delete task - gorm creates TIMESTAMP deleted_at, not actual delete
func (task Task) Delete(ID uint) []error {
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return []error{err}
	}

	errors = db.First(&task, "id = ?", ID).GetErrors()
	if len(errors) != 0 {
		return errors
	}
	errors = db.Delete(&task).GetErrors()
	return errors
}

//fetch given task by ID
func (task Task) Fetch(ID uint) (TaskEntity, []error) {
	var tempTask TaskEntity
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return tempTask, []error{err}
	}

	errors = db.First(&task, "id = ? and isnull(deleted_at)", ID).Scan(&tempTask).GetErrors()
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
