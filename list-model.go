package main

import "github.com/jinzhu/gorm"

//TodoList is a model used by gorm for todo_lists table.
type ListModel struct {
	gorm.Model
	Name  string      `sql:"not null;unique"`
	Tasks []TaskModel `gorm:"ForeignKey:LID"`
}
