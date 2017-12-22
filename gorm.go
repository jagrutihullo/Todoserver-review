package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type List struct {
	//gorm.Model
	Model     gorm.Model `gorm:"embedded"`
	name      string     `gorm:"primary_key"`
	createdAt time.Time
}

/*
func main() {
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("error 1")
		panic(err.Error())
	}
	defer db.Close()

	//firstList := List{}
	//db.First(&firstList)
	if db.HasTable("list") {
		allLists := List{}
		db.Find(&allLists)
		fmt.Println(allLists)
	} else {
		fmt.Println("Table does not exists")
	}
}
*/
