package database

import (
	"backend/model"
	"errors"
)

func SaveTodo(todo *model.Todo) error {
	query := `INSERT INTO todo (title, description, completed) VALUES (?,?,?)`
	err := ExecuteCommand(DB, query, todo.Title, todo.Description, todo.Completed)
	return err
}

var ErrTodoNotFound = errors.New("todo not found")

// Return empty slice if none
func AllTodos() ([]model.Todo, error) {
	query := `SELECT * FROM todo`
	rows, err := DB.Query(query)
	res := make([]model.Todo, 0)
	//Early return
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var todo model.Todo
		rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		res = append(res, todo)
	}
	return res, nil
}

func TodosByPropertyAndValue(property string, value any) ([]model.Todo, error) {
	query := "SELECT * FROM todo WHERE " + property + " = ?"
	rows, err := DB.Query(query, value)
	res := make([]model.Todo, 0)
	//Early return
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var todo model.Todo
		rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		res = append(res, todo)
	}
	return res, nil
}

func CompletedTodos() ([]model.Todo, error) {
	return TodosByPropertyAndValue("completed", true)
}

func UncompletedTodos() ([]model.Todo, error) {
	return TodosByPropertyAndValue("completed", false)
}

func UpdateTodo(todo *model.Todo) error {
	var count int
	countQuery := "SELECT COUNT(*) FROM todo WHERE id = ?"
	err := DB.QueryRow(countQuery, todo.ID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		// El "todo" no existe, devolver un error personalizado
		return ErrTodoNotFound
	}
	query := "UPDATE todo SET title = ?, description = ?, completed = ? WHERE id = ?"
	_, err = DB.Exec(query, todo.Title, todo.Description, todo.Completed, todo.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodoById(todo *model.Todo, id int) error {
	var count int
	countQuery := "SELECT COUNT(*) FROM todo WHERE id = ?"
	err := DB.QueryRow(countQuery, todo.ID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		// El "todo" no existe, devolver un error personalizado
		return ErrTodoNotFound
	}
	query := "UPDATE todo SET title = ?, description = ?, completed = ? WHERE id = ?"
	_, err = DB.Exec(query, todo.Title, todo.Description, todo.Completed, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todo *model.Todo) error {
	var count int
	countQuery := "SELECT COUNT(*) FROM todo WHERE id = ?"
	err := DB.QueryRow(countQuery, todo.ID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		// El "todo" no existe, devolver un error personalizado
		return ErrTodoNotFound
	}
	query := "DELETE FROM todo WHERE id = ?"
	_, err = DB.Exec(query, todo.ID)
	return err
}

func DeleteTodoById(id int64) error {
	var count int
	countQuery := "SELECT COUNT(*) FROM todo WHERE id = ?"
	err := DB.QueryRow(countQuery, id).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		// El "todo" no existe, devolver un error personalizado
		return ErrTodoNotFound
	}
	query := "DELETE FROM todo WHERE id = ?"
	_, err = DB.Exec(query, id)
	return err
}
