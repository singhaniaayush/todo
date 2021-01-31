package service

import (
	"todo/server/models"
	"todo/server/utils"
)

// RetrieveAllTodos - fetch all todos
func RetrieveAllTodos() ([]models.Todos, error) {
	var todos []models.Todos
	err := utils.Db.Find(&todos).Error

	if err != nil {
		return nil, err
	}
	return todos, nil
}

// RetrieveTodoByID - fetch todo by id
func RetrieveTodoByID(id string) (models.Todos, error) {
	var todo models.Todos
	err := utils.Db.First(&todo, id).Error

	if err != nil {
		return models.Todos{}, err
	}
	return todo, nil
}

// DeleteTodoByID - fetch todo by id
func DeleteTodoByID(id string) error {
	var todo models.Todos
	err := utils.Db.First(&todo, id).Error

	if err != nil {
		return err
	}

	utils.Db.Delete(&todo)
	return nil
}

// SaveTodo - save todo
func SaveTodo(todo *models.Todos) error {
	err := utils.Db.Create(&todo).Error
	return err
}

// UpdateTodo - update todo if present
func UpdateTodo(todo *models.Todos, todoID string) error {
	var toUpdate models.Todos
	err := utils.Db.First(&toUpdate, todoID).Error

	if err != nil {
		return err
	}

	utils.Db.Model(&toUpdate).Updates(todo)
	return nil
}
