package models

import (
	"net/http"

	"github.com/muasx/todo_api/db"
)

type Todo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete int    `json:"complete"`
}

func GetTodos() (Response, error) {
	var todo Todo
	var todos []Todo
	var res Response

	con := db.CreateCon()

	rows, err := con.Query("SELECT id, name, complete FROM todos")
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		if err = rows.Scan(&todo.ID, &todo.Name, &todo.Complete); err != nil {
			return res, err
		}

		todos = append(todos, todo)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = todos

	return res, nil
}

func Store(name string, complete int) (Response, error) {
	var res Response

	con := db.CreateCon()

	stmt, err := con.Prepare("INSERT INTO todos (name, complete) VALUES (?, ?)")
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, complete)
	if err != nil {
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Success Insert Todo"
	res.Data = lastInsertedID

	return res, nil
}

func Update(id int, name string, complete int) (Response, error) {
	var res Response

	con := db.CreateCon()

	stmt, err := con.Prepare("UPDATE todos SET name = ?, complete = ? WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, complete, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Updated"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func Delete(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	stmt, err := con.Prepare("DELETE FROM todos WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Todo Deleted"
	res.Data = map[string]int64{
		"row_affected": rowsAffected,
	}

	return res, nil
}
