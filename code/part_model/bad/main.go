package main

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Integral int	`db:"integral"`
}
var db = *sqlx.DB
func init () {
	newDB, err:= sqlx.Connect("mysql", "")
	if err != nil {panic(err)}
	db = newDB
}
// 不好的示例
func ListPartUser() (userList []User) {
	db.Select(&userList, `SELECT id, name FROM user`)
	return userList
}
func main() {
	userList := ListPartUser()
	// age 是 zero value
	log.Print(userList[0].Age)
}
