package main

import "github.com/jinzhu/gorm"

//todolist gorm model used by gorm
type TodoList struct {
	gorm.Model
	Name  string      `sql:"not null;unique"`
	Tasks []TaskModel `gorm:"ForeignKey:LID"`
}
