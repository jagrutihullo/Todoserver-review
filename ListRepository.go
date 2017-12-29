package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//todolist gorm model used by gorm
type TodoList struct {
	gorm.Model
	Name  string      `sql:"not null;unique"`
	Tasks []TaskModel `gorm:"ForeignKey:LID"`
}

//repository for todolist
type TodoListRepository interface {
	Create(todoEntity List) []error

	Fetch(name string) (List, []error)

	Update(todoEntity List) []error

	Delete(name string) []error

	FetchAll() ([]List, []error)
}

//create list with name
func (todoList TodoList) Create(todoEntity List) []error {
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return []error{err}
	}

	if db.HasTable(&TodoList{}) == false {
		db.CreateTable(&TodoList{})
		db.CreateTable(&TaskModel{})
	}

	todoList.Name = todoEntity.Name
	errors = db.Create(&todoList).GetErrors()
	return errors
}

//update list name only
func (todoList TodoList) Update(todoEntity List) []error {
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return []error{err}
	}

	errors = db.Find(&todoList, todoEntity.ID).GetErrors()
	if len(errors) != 0 {
		return errors
	}
	todoList.Name = todoEntity.Name
	errors = db.Save(&todoList).GetErrors()
	return errors
}

//delete list - gorm creates TIMESTAMP deleted_at, not actual delete
func (todoList TodoList) Delete(name string) []error {
	var tasks []TaskModel
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return []error{err}
	}

	errors = db.First(&todoList, "name = ? and isnull(deleted_at)", name).GetErrors()
	if len(errors) != 0 {
		return errors
	}

	db.Find(&tasks, "l_id = ? and isnull(deleted_at)", todoList.ID)
	for i := range tasks {
		errors = db.Delete(tasks[i]).GetErrors()
		if len(errors) != 0 {
			return errors
		}
	}

	errors = db.Delete(todoList).GetErrors()
	return errors
}

//fetch given list by name and tasks under that list
func (todoList TodoList) Fetch(name string) (List, []error) {
	var tempList List
	var errors []error
	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return tempList, []error{err}
	}

	errors = db.First(&todoList, "name = ? and isnull(deleted_at)", name).Scan(&tempList).GetErrors()
	if len(errors) != 0 {
		return tempList, errors
	}

	errors = db.Find(&todoList.Tasks, "l_id = ? and isnull(deleted_at)", todoList.ID).Scan(&tempList.Tasks).GetErrors()
	return tempList, errors
}

//fetch all lists without showing tasks
func (todoList TodoList) FetchAll() ([]List, []error) {
	var tempList []List
	var errors []error

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return tempList, []error{err}
	}

	errors = db.Table("todo_lists").Where("isnull(deleted_at)").Select("id, name, created_at, updated_at").Scan(&tempList).GetErrors()
	return tempList, errors

}

/*func main() {
	//var t []TodoListEntity
	var t1 TodoListEntity
	var task []TaskEntity
	var err []error
	var tModel TodoList
	// t1 = TodoListEntity{
	// 	Name: "grocery list",
	// }

	// err = tModel.Create(t1)
	// fmt.Println(err)

	// t1 = TodoListEntity{
	// 	Name: "grocery list 2",
	// }
	// err = tModel.Create(t1)
	// fmt.Println(err)

	// t1 = TodoListEntity{
	// 	Name: "grocery list 3",
	// }
	// err = tModel.Create(t1)
	// fmt.Println(err)

	// t1 = TodoListEntity{
	// 	ID:   1,
	// 	Name: "grocery list december",
	// }
	// err = tModel.Update(t1)
	// fmt.Println(len(err))

	// t1, task, err = tModel.Fetch("grocery list december")
	// fmt.Println(t)
	// fmt.Println(task)
	// fmt.Println(err)

	t1, task, err = tModel.Fetch("grocery list 2")
	fmt.Println(t1)
	fmt.Println(task)
	fmt.Println(err)

	// err = tModel.Delete("grocery list december")
	// fmt.Println(err)

	// t, err = tModel.FetchAll()
	// fmt.Println(t)
	// fmt.Println(err)
}*/
