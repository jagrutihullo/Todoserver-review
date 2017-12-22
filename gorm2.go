package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	ID        int
	Firstname string
	Lastname  string
	Salary    float32
}

/*
func main() {

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("error ", err)
	}

	db.CreateTable(&User{})

	user := User{
		ID:        101,
		Lastname:  "Dent",
		Firstname: "Adent",
		Salary:    5000,
	}

	if db.HasTable(&User{}) {
		db.Create(&user)
		fmt.Println("row created")
	} else {
		fmt.Println("row not created")
	}

	allUsers := User{}
	db.Find(&allUsers)
	fmt.Println(allUsers)

}
*/
