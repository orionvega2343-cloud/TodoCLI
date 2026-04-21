package cmd

import (
	"TodoCLI/db"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ValidateCmd(command string, conn *sql.DB) error {
	switch command {
	case "add":
		task := os.Args[2]
		err := addTask(conn, task)
		if err != nil {
			return err
		}
	case "list":
		err := listTask(conn)
		if err != nil {
			return err
		}
	case "done":
		task := os.Args[3]
		id := os.Args[2]
		parsed, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		err = editTask(conn, task, parsed)
		if err != nil {
			return err
		}
	case "delete":
		id := os.Args[2]
		parsed, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		err = deleteTask(conn, parsed)
		if err != nil {
			return err
		}
	}
	return nil
}

func addTask(conn *sql.DB, task string) error {
	err := db.Create(conn, task)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func listTask(conn *sql.DB) error {
	tasks, err := db.GetAll(conn)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		fmt.Println(task)
	}
	return nil
}

func editTask(conn *sql.DB, task string, id int) error {
	err := db.Update(conn, task, id)
	if err != nil {
		return err
	}
	return nil
}

func deleteTask(conn *sql.DB, id int) error {
	err := db.Delete(conn, id)
	if err != nil {
		return err
	}
	return nil
}
