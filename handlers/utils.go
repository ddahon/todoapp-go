package handlers

import (
	"database/sql"
	"log"

	"github.com/ddahon/go-todoapp/models"
)

func getData(db *sql.DB) []models.Todo {
	var todos []models.Todo
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var t models.Todo
		rows.Scan(&t.Description, &t.Id)
		todos = append(todos, t)
	}

	return todos
}
