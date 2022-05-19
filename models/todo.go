package models

import (
	"encoding/json"
	"errors"
	validator "gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

var validation = validator.New()

type Todo struct {
	ID     uint    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name   *string `json:"name" validate:"min=3,max=100" gorm:"not null"`
	Status bool    `json:"status"`
}

type todoResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}

func init() {
	if !db.Migrator().HasTable(&Todo{}) {
		db.Migrator().CreateTable(&Todo{})
	}
}

func GetAllTodo() (data []Todo, err error) {
	var todoList []Todo
	result := db.Find(&todoList)
	if result.Error != nil {
		return todoList, result.Error
	}
	return todoList, nil
}

func GetTodoById(id string) (data Todo, err error) {
	todoData := Todo{}
	errorMessage := db.First(&todoData, id).Error
	if errors.Is(errorMessage, gorm.ErrRecordNotFound) {
		return todoData, errorMessage
	}
	return todoData, nil
}

func CreateTodo(todo Todo) (data Todo, err error) {
	errorMessage := validation.Struct(&todo)
	if errorMessage != nil {
		return todo, errorMessage
	}
	results := db.Create(&todo)
	if results.Error != nil {
		return todo, results.Error
	}
	return todo, nil
}

func UpdateTodo(id string, dataTodo []byte) (data Todo, err error) {
	todoData := Todo{}
	errorMessage := db.First(&todoData, id).Error
	if errorMessage != nil {
		return todoData, errorMessage
	}
	json.Unmarshal(dataTodo, &todoData)
	errorMessage = validation.Struct(&todoData)
	if errorMessage != nil {
		return todoData, errorMessage
	}
	result := db.Save(&todoData)
	if result.Error != nil {
		return todoData, result.Error
	}
	return todoData, nil
}

func DeleteTodo(id string) (idTodo string, err error) {
	result := db.Delete(&Todo{}, id)
	if result.Error != nil {
		return idTodo, result.Error
	}
	return idTodo, nil
}

func CustomResponseTodo(statusCode int, data interface{}, message string) (responseMessage todoResponse) {
	return todoResponse{statusCode, data, message}
}
