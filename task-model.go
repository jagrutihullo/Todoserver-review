package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

//TaskModel is a model used by gorm for task_models table.
type TaskModel struct {
	gorm.Model
	Description string
	Status      string
	Deadline    time.Time
	LID         uint
}

//TaskToModel converts task entity to task model
func TaskToModel(task Task, taskModel TaskModel) TaskModel {
	taskModel.Description = task.Description
	taskModel.Status = task.Status
	taskModel.Deadline = task.Deadline
	return taskModel
}
