package db

import (
	"TodoCLI/models"
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
    ID SERIAL PRIMARY KEY,
    task TEXT NOT NULL
)`)
	if err != nil {
		return err
	}
	return nil
}

func Create(db *sql.DB, task string) error {
	_, err := db.Exec(`INSERT INTO tasks (task) VALUES ($1)`, task)
	if err != nil {
		return err
	}
	return nil
}

func Update(db *sql.DB, task string, id int) error {
	_, err := db.Exec(`UPDATE tasks SET task = $1 WHERE id = $2 `, task, id)
	if err != nil {
		return err
	}
	return nil
}

func GetAll(db *sql.DB) ([]models.Task, error) {
	rows, err := db.Query(`SELECT id,task FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)

	}
	return tasks, nil
}

func Delete(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
