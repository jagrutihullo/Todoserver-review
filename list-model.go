package main

import "github.com/jinzhu/gorm"

//TodoList is a model used by gorm for todo_lists table.
type TodoList struct {
	gorm.Model
	Name  string      `sql:"not null;unique"`
	Tasks []TaskModel `gorm:"ForeignKey:LID"`
}
