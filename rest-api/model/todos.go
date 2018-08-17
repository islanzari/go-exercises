package model

import (
	"database/sql"
	"fmt"
)

type TodosModel struct {
	DB *sql.DB
}
type Todo struct {
	ID          int
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

func (tm *TodosModel) UpdateToDo(id int, name string, description string) error {

	sqlStatement := `UPDATE todos SET name = $1, description = $2, updated_at = now() where id = $3`
	_, err := tm.DB.Exec(sqlStatement, name, description, id)
	return err
}

func (tm *TodosModel) FetchToDo(id int) (Todo, error) {
	todo := Todo{}
	sqlStatement := `SELECT id, name, description, created_at, updated_at FROM todos where id = $1 `
	row := tm.DB.QueryRow(sqlStatement, id)
	err := row.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)

	return todo, err
}

func (tm *TodosModel) DeleteToDo(id int) error {
	sqlStatement := ("DELETE FROM todos WHERE id = $1")
	_, err := tm.DB.Exec(sqlStatement, id)

	return err
}

func (tm *TodosModel) FetchToDos() ([]Todo, error) {
	todos := []Todo{}
	sqlStatement := `SELECT id, name, description, created_at, updated_at FROM todos`
	rows, err := tm.DB.Query(sqlStatement)
	if err != nil {
		return todos, err
	}

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()

}

func (tm *TodosModel) CreateToDo(name string, description string) error {

	_, err := tm.DB.Exec("INSERT INTO todos(name, description) VALUES($1, $2)", name, description)
	if err != nil {
		fmt.Println("ERROR saving to db - ", err)
	}
	return err
}
