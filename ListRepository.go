package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//gorm model used by gorm
type TodoList struct {
	gorm.Model
	Name  string `sql:"not null;unique"`
	Tasks []Task `gorm:"ForeignKey:LID"`
}

//repository for todolist
type TodoListRepository interface {
	Create(todoEntity TodoListEntity) []error

	Fetch(name string) (TodoListEntity, []TaskEntity, error)

	Update(todoEntity TodoListEntity) []error

	Delete(name string) []error

	FetchAll() ([]TodoListEntity, []error)
}

//create list with name
func (todoList TodoList) Create(todoEntity TodoListEntity) []error {
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return []error{err}
	}
	fmt.Println(err)

	if db.HasTable(&TodoList{}) == false {
		db.CreateTable(&TodoList{})
		db.CreateTable(&Task{})
	}

	todoList.Name = todoEntity.Name
	errors = db.Create(&todoList).GetErrors()
	return errors
}

//update list name only
func (todoList TodoList) Update(todoEntity TodoListEntity) []error {
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return []error{err}
	}

	errors = db.Find(&todoList, todoEntity.ID).GetErrors()
	if len(errors) != 0 {
		fmt.Println(errors)
		return errors
	}
	todoList.Name = todoEntity.Name
	errors = db.Save(&todoList).GetErrors()
	return errors
}

//delete list - gorm creates TIMESTAMP deleted_at, not actual delete
func (todoList TodoList) Delete(name string) []error {
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return []error{err}
	}

	errors = db.First(&todoList, "name = ?", name).GetErrors()
	if len(errors) != 0 {
		return errors
	}
	errors = db.Delete(&todoList).GetErrors()
	return errors
}

//fetch given list by name and tasks under that list
func (todoList TodoList) Fetch(name string) (TodoListEntity, []TaskEntity, []error) {
	var tempList TodoListEntity
	var taskList []Task
	var tempTasks []TaskEntity
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return tempList, tempTasks, []error{err}
	}

	errors = db.First(&todoList, "name = ? and isnull(deleted_at)", name).Scan(&tempList).GetErrors()
	if len(errors) != 0 {
		return tempList, tempTasks, errors
	}
	errors = db.Find(&taskList, "l_id = ?", todoList.ID).Scan(&tempTasks).GetErrors()
	if len(errors) != 0 {
		return tempList, tempTasks, errors
	}
	return tempList, tempTasks, errors
}

//fetch all lists without showing tasks
func (todoList TodoList) FetchAll() ([]TodoListEntity, []error) {
	var tempList []TodoListEntity
	var errors []error

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return tempList, []error{err}
	}

	errors = db.Table("todo_lists").Where("isnull(deleted_at)").Select("id, name, created_at, updated_at").Scan(&tempList).GetErrors()
	return tempList, errors

}

/* func main() {
	var t []TodoListEntity
	var t1 TodoListEntity
	var task []TaskEntity
	var err []error
	var tModel TodoList
	t1 = TodoListEntity{
		Name: "grocery list",
	}

	err = tModel.Create(t1)
	fmt.Println(err)

	t1 = TodoListEntity{
		Name: "grocery list 2",
	}
	err = tModel.Create(t1)
	fmt.Println(err)

	t1 = TodoListEntity{
		Name: "grocery list 3",
	}
	err = tModel.Create(t1)
	fmt.Println(err)

	t1 = TodoListEntity{
		ID:   1,
		Name: "grocery list december",
	}
	err = tModel.Update(t1)
	fmt.Println(len(err))

	t1, task, err = tModel.Fetch("grocery list december")
	fmt.Println(t)
	fmt.Println(task)
	fmt.Println(err)

	t1, task, err = tModel.Fetch("grocery list 2")
	fmt.Println(t)
	fmt.Println(task)
	fmt.Println(err)

	err = tModel.Delete("grocery list 2")
	fmt.Println(err)

	t, err = tModel.FetchAll()
	fmt.Println(t)
	fmt.Println(err)
}*/
