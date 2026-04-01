package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Person struct {
	UserId   string `db:"id"`
	Username string `db:"name"`
	Age      int    `db:"age"`
	Address  string `db:"address"`
}

func init() {
	conn, err := sqlx.Open("mysql", "root:Zhaook123!@tcp(127.0.0.1:3306)/data")
	if err != nil {
		fmt.Println("Open mysql failed", err)
		return
	}
	db = conn
}
func main() {
	list()
	defer db.Close()
}
func list() {
	var persons []Person
	err := db.Select(&persons, "select * from user")
	if err != nil {
		fmt.Println("list err", err)
		return
	}
	for _, person := range persons {
		fmt.Printf("list succ,%+v\n", person)
	}
}
