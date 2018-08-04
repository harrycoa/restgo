package main

import (
	"fmt"

	"./models"
)

func main() {
	models.CreateConnection()
	models.Ping()
	//result := models.ExistTable("users")
	//fmt.Println(result)
	//models.CreateTables()

	user := models.CreateUser("aaaaaa", "aaa123", "a@a")
	fmt.Println(user)

	user.Username = "cambio de nombre"
	user.Save()

	models.Closeconnection()
}
