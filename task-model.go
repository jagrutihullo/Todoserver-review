package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

//task gorm model used by gorm
type TaskModel struct {
	gorm.Model
	Description string
	Status      string
	Deadline    time.Time
	LID         uint
}

//function to convert task entity to task model
func TaskToModel(task Task) TaskModel {
	var taskModel TaskModel
	taskModel.Description = task.Description
	taskModel.Status = task.Status
	taskModel.Deadline = task.Deadline
	return taskModel
}
