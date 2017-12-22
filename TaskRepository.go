package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//exported
type Task struct {
	gorm.Model
	Description string
	Status      string
	Deadline    time.Time
	LID         uint
}
