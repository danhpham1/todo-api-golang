package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"todo-api/models"
)

type TodoController struct {
	beego.Controller
}

func (todo *TodoController) GetAll() {
	todoList, err := models.GetAllTodo()
	if err != nil {
		todo.Data["json"] = models.CustomResponseTodo(500, []models.Todo{}, err.Error())
		todo.ServeJSON()
		return
	}
	dataResponse := models.CustomResponseTodo(200, todoList, "Lấy danh sách thành công")
	todo.Data["json"] = &dataResponse
	todo.ServeJSON()
	return
}

func (todo *TodoController) Post() {
	todoData := models.Todo{}
	json.Unmarshal(todo.Ctx.Input.RequestBody, &todoData)
	data, err := models.CreateTodo(todoData)
	if err != nil {
		todo.Data["json"] = models.CustomResponseTodo(500, data, err.Error())
		todo.ServeJSON()
		return
	}
	todo.Data["json"] = models.CustomResponseTodo(201, data, "Tạo thành công")
	todo.ServeJSON()
	return
}

func (todo *TodoController) GetTodoById() {
	id := todo.Ctx.Input.Param(":id")
	todoData, err := models.GetTodoById(id)
	if err != nil {
		todo.Data["json"] = models.CustomResponseTodo(500, nil, err.Error())
		todo.ServeJSON()
		return
	}
	todo.Data["json"] = models.CustomResponseTodo(200, todoData, "Lấy thành công")
	todo.ServeJSON()
	return
}

func (todo *TodoController) Put() {
	id := todo.Ctx.Input.Param(":id")
	if id != "" {
		todoData, err := models.UpdateTodo(id, todo.Ctx.Input.RequestBody)
		if err != nil {
			todo.Data["json"] = models.CustomResponseTodo(500, todoData, err.Error())
			todo.ServeJSON()
			return
		}
		todo.Data["json"] = models.CustomResponseTodo(200, todoData, "Cập nhập thành công")
		todo.ServeJSON()
		return
	}
}

func (todo *TodoController) Delete() {
	id := todo.Ctx.Input.Param(":id")
	_, err := models.DeleteTodo(id)
	if err != nil {
		todo.Data["json"] = models.CustomResponseTodo(500, id, err.Error())
		todo.ServeJSON()
		return
	}
	todo.Data["json"] = models.CustomResponseTodo(200, id, "Xóa thành công")
	todo.ServeJSON()
	return
}
