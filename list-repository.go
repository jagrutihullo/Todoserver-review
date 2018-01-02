package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//repository for todolist
type TodoListRepository interface {
	Create(todoEntity List) error

	Fetch(name string) (List, error)

	Update(todoEntity List) error

	Delete(name string) error

	FetchAll() ([]List, error)
}

type GormListRepo struct {
}

//create list with name
func (glr GormListRepo) Create(todoEntity List) error {
	var error1 error
	var errorsArr []error
	var todoList TodoList

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	if db.HasTable(&TodoList{}) == false {
		db.CreateTable(&TodoList{})
		db.CreateTable(&TaskModel{})
	}

	todoList.Name = todoEntity.Name
	errorsArr = db.Create(&todoList).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1

}

//update list name only
func (glr GormListRepo) Update(todoEntity List) error {
	var error1 error
	var errorsArr []error
	var todoList TodoList

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	errorsArr = db.Find(&todoList, "id = ? and isnull(deleted_at)", todoEntity.ID).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return error1
	}
	todoList.Name = todoEntity.Name
	errorsArr = db.Save(&todoList).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//delete list - gorm creates TIMESTAMP deleted_at, not actual delete
func (glr GormListRepo) Delete(name string) error {
	var tasks []TaskModel
	var error1 error
	var errorsArr []error
	var todoList TodoList

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	errorsArr = db.First(&todoList, "name = ? and isnull(deleted_at)", name).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return error1
	}

	db.Find(&tasks, "l_id = ? and isnull(deleted_at)", todoList.ID)
	for i := range tasks {
		errorsArr = db.Delete(tasks[i]).GetErrors()
		if len(errorsArr) != 0 {
			error1 = ErrorsConv(errorsArr)
			return error1
		}
	}

	errorsArr = db.Delete(todoList).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return error1
}

//fetch given list by name and tasks under that list
func (glr GormListRepo) Fetch(name string) (List, error) {
	var tempList List
	var error1 error
	var errorsArr []error
	var todoList TodoList

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return tempList, err
	}

	errorsArr = db.First(&todoList, "name = ? and isnull(deleted_at)", name).Scan(&tempList).GetErrors()
	if len(errorsArr) != 0 {
		error1 = ErrorsConv(errorsArr)
		return tempList, error1
	}

	errorsArr = db.Find(&todoList.Tasks, "l_id = ? and "+
		"isnull(deleted_at)", todoList.ID).Scan(&tempList.Tasks).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return tempList, error1
}

//fetch all lists without showing tasks
func (glr GormListRepo) FetchAll() ([]List, error) {
	var tempList []List
	var error1 error
	var errorsArr []error

	db, err := gorm.Open("mysql", "root:root@/todoserver?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return tempList, err
	}

	errorsArr = db.Table("todo_lists").Where("isnull(deleted_at)").Select("id, name, " +
		"created_at, updated_at").Scan(&tempList).GetErrors()
	error1 = ErrorsConv(errorsArr)
	return tempList, error1

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
