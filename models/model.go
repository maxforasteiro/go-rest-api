package models

import (
	"../config"
	"fmt"
	_ "github.com/lib/pq"
)

func GetTodos(todo *[]Todo) (err error) {
	if err := config.DB.Find(todo).Error; err != nil {
		fmt.Println("get all")
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error) {
	if err := config.DB.Create(todo).Error; err != nil {
		fmt.Println("create todo")
		return err
	}
	return nil
}

func GetTodo(todo *Todo, id string) (err error) {
	fmt.Println(id)
	if err := config.DB.First(&todo, "id = ?", id).Error; err != nil {
		fmt.Println("get todo")
		return err
	}
	return nil
}

func UpdateTodo(todo *Todo) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

func DeleteTodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Where("id = ?", id).Delete(&todo)
	return nil
}
