package main

import (
	"TodoCLI/cmd"
	"TodoCLI/db"
	"fmt"
	"log"
	"os"
)

func main() {
	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=todos sslmode=disable",
		os.Getenv("DB_PASSWORD"),
	)
	conn, err := db.NewDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.CreateTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Введите команду add,list,done,delete")
		return
	}

	command := os.Args[1]
	err = cmd.ValidateCmd(command, conn)
	if err != nil {
		log.Fatal(err)
	}

}
