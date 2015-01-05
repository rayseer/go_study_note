package model

import (
	//类似于PHP里的PDO，多数都用这个来连接数据库
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Model struct {
	Db *sql.DB
}

func Connect() *Model {
	fmt.Println("Connect to DB...")
	db, err := sql.Open("mysql", "root:000111@tcp(localhost:3306)/golangtest")
	if err != nil {
		log.Fatal(err)
	}
	return &Model{db}
}

func (model *Model) FindOne() {
	/*
		id := 1
		rows, err := model.Query("SELECT * FROM users WHERE id = ?", id)
		if err != nil {
			log.Fatal(err)
		}
		return rows
	*/
}

func FindAll() {

}

func FindCount() {

}

func Insert() {

}

func Update() {

}

func Delete() {

}
