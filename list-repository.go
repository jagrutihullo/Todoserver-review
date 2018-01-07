package main

import (
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//ListRepository is repository interface for list and task functions
type ListRepository interface {
	Create(listEntity List) error

	FetchByID(list List) (List, error)

	Update(listEntity List) error

	Delete(list List) error

	FetchAll() ([]List, error)

	CreateTask(list List) error

	FetchTaskByID(list List) (List, error)

	UpdateTask(list List) error

	DeleteTask(list List) error
}

func dbConnection() (*gorm.DB, error) {

	DBHost := os.Getenv("DBHost")
	Username := os.Getenv("Username")
	Password := os.Getenv("Password")
	Host := os.Getenv("Host")
	DBName := os.Getenv("DBName")
	Charset := os.Getenv("Charset")
	ParseTime := os.Getenv("ParseTime")
	Loc := os.Getenv("Loc")

	db, dbError := gorm.Open(DBHost, Username+"@"+Host+":"+Password+"@/"+DBName+"?charset="+
		Charset+"&parseTime="+ParseTime+"&loc="+Loc)
	return db, dbError
}

//GormListRepo is a structure that implements TodoListRepo functions
type GormListRepo struct {
}

//Create is a Gorm function to create list
func (glr GormListRepo) Create(listEntity List) error {
	var (
		dbError  error
		todoList ListModel
	)

	db, dbError := dbConnection()
	if dbError != nil {
		return dbError
	}
	defer db.Close()

	if db.HasTable(&ListModel{}) == false {
		db.CreateTable(&ListModel{})
		db.CreateTable(&ListModel{})
	}

	todoList.Name = listEntity.Name
	dbError = singleError(db.Create(&todoList).GetErrors())
	return dbError
}

//Update is a Gorm function to update list name
func (glr GormListRepo) Update(listEntity List) error {
	var (
		dbError  error
		todoList ListModel
	)

	db, dbError := dbConnection()
	if dbError != nil {
		return dbError
	}
	defer db.Close()

	dbError = singleError(db.Find(&todoList, "id = ? and "+
		"isnull(deleted_at)", listEntity.ID).GetErrors())
	if dbError != nil {
		return dbError
	}

	todoList.Name = listEntity.Name
	dbError = singleError(db.Save(&todoList).GetErrors())
	return dbError
}

//Delete is a Gorm function to delete list and tasks under it
//Gorm creates deleted_at TIMESTAMP, it does not actually deletes record
func (glr GormListRepo) Delete(list List) error {
	var (
		dbError  error
		todoList ListModel
	)

	db, dbError := dbConnection()
	if dbError != nil {
		return dbError
	}
	defer db.Close()

	dbError = singleError(db.First(&todoList, "id = ? and "+
		"isnull(deleted_at)", list.ID).GetErrors())
	if dbError != nil {
		return dbError
	}

	dbError = singleError(db.Table("task_models").Where("l_id = ? and "+
		"isnull(deleted_at)", todoList.ID).Delete(&TaskModel{}).GetErrors())
	if dbError != nil {
		return dbError
	}

	dbError = singleError(db.Delete(todoList).GetErrors())
	return dbError
}

//FetchByID is a Gorm function to access list by name and tasks under it
func (glr GormListRepo) FetchByID(list List) (List, error) {
	var (
		tempList  List
		dbError   error
		todoList  ListModel
		taskModel TaskModel
	)

	db, dbError := dbConnection()
	if dbError != nil {
		return tempList, dbError
	}
	defer db.Close()

	dbError = singleError(db.First(&todoList, "id = ? and "+
		"isnull(deleted_at) ", list.ID).Scan(&tempList).GetErrors())
	if dbError != nil {
		return tempList, dbError
	}

	db.Find(&taskModel, "l_id = ? and "+
		"isnull(deleted_at)", tempList.ID).Scan(&tempList.Tasks)
	return tempList, dbError
}

//FetchAll is a Gorm function to access all lists
func (glr GormListRepo) FetchAll() ([]List, error) {
	var (
		tempList  []List
		taskModel TaskModel
		dbError   error
	)

	db, dbError := dbConnection()
	if dbError != nil {
		return tempList, dbError
	}
	defer db.Close()

	dbError = singleError(db.Table("list_models").Where("isnull(deleted_at)").Select("id, " +
		"name, created_at, updated_at").Scan(&tempList).GetErrors())

	if dbError == nil {
		for i := range tempList {
			db.Find(&taskModel, "l_id = ? and "+
				"isnull(deleted_at)", tempList[i].ID).Scan(&tempList[i].Tasks)
		}
	}
	return tempList, dbError
}

//CreateTask is a Gorm function to create task under a list
func (glr GormListRepo) CreateTask(list List) error {
	var (
		dbError   error
		taskModel TaskModel
		tempList  List
	)

	db, dbError := dbConnection()
	if dbError != nil {
		return dbError
	}
	defer db.Close()

	tempList, dbError = glr.FetchByID(list)
	if dbError != nil {
		return dbError
	}
	taskModel.TaskToModel(list.Tasks[0])
	taskModel.LID = tempList.ID
	dbError = singleError(db.Create(&taskModel).GetErrors())
	return dbError
}

//UpdateTask is a Gorm function to update task
func (glr GormListRepo) UpdateTask(list List) error {
	var (
		dbError   error
		taskModel TaskModel
	)

	db, dbError := dbConnection()
	if dbError != nil {
		return dbError
	}
	defer db.Close()

	dbError = singleError(db.First(&taskModel, "id = ? and "+
		"isnull(deleted_at)", list.Tasks[0].ID).GetErrors())
	if dbError != nil {
		return dbError
	}

	taskModel.TaskToModel(list.Tasks[0])
	dbError = singleError(db.Save(&taskModel).GetErrors())
	return dbError
}

//DeleteTask is a Gorm function to delete task
//Gorm creates deleted_at TIMESTAMP, it does not actually deletes record
func (glr GormListRepo) DeleteTask(list List) error {
	var (
		dbError   error
		taskModel TaskModel
	)

	db, dbError := dbConnection()
	if dbError != nil {
		return dbError
	}
	defer db.Close()

	dbError = singleError(db.First(&taskModel, "id = ? and "+
		"isnull(deleted_at)", list.Tasks[0].ID).GetErrors())
	if dbError != nil {
		return dbError
	}

	dbError = singleError(db.Delete(&taskModel).GetErrors())
	return dbError
}

//FetchTaskByID is a Gorm function to access task by id
func (glr GormListRepo) FetchTaskByID(list List) (List, error) {
	var (
		tempList  List
		dbError   error
		taskModel TaskModel
	)

	db, dbError := dbConnection()
	if dbError != nil {
		return tempList, dbError
	}
	defer db.Close()

	tempList.Tasks = make([]Task, 1)

	dbError = singleError(db.First(&taskModel, "id = ? and "+
		"isnull(deleted_at)", list.Tasks[0].ID).Scan(&tempList.Tasks[0]).GetErrors())
	tempList.ID = taskModel.LID
	list, dbError = glr.FetchByID(tempList)
	if dbError != nil {
		return list, dbError
	}

	tempList.Name = list.Name
	tempList.CreatedAt = list.CreatedAt
	tempList.UpdatedAt = list.UpdatedAt
	return tempList, dbError
}

//singleError converts array of errors into single error
func singleError(errorsArray []error) error {
	var dbError error
	var errString string

	for i := range errorsArray {
		errString += errorsArray[i].Error()
	}
	dbError = errors.New(errString)
	return dbError
}
