package main

import (
	"log"
)

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
type PartUser_ID_Name struct {
	ID string
	Name string
}
var db = *sqlx.DB
func init () {
	newDB, err:= sqlx.Connect("mysql", "")
	if err != nil {panic(err)}
	db = newDB
}
func ListPartUser_ID_Name() (partUserList []PartUser_ID_Name) {
	userList := []User{}
	db.Select(&userList, `SELECT id, name FROM user`)
	for _, user := range userList {
		partUserList = append(partUserList, PartUser_ID_Name{
			ID: user.ID,
			Name: user.Name,
		})
	}
	return partUserList
}
// 单个字段直接使用 ListUser{字段} 并返回 []字段类型
func ListUserID() (userIDList []string) {
	userList := []User{}
	db.Select(&userIDList, `SELECT id FROM user`)
	for _, user := range userList {
		userIDList = append(userIDList, user.ID)
	}
	return userIDList
}
func main() {
	partUserList := ListPartUser_ID_Name()
	log.Print(partUserList[0].Name)
	// 此时使用 partUserList[0].Age 会报错，因为 PartUser_ID_Name 中没有 Age
}

