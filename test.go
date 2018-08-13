package main

import (
	"os"
	"fmt"
	"./config"
	"./models"
)

func main() {
	models.CreateConnection()
	//models.Ping()
	//result := models.ExistTable("users")
	//fmt.Println(result)
	//models.CreateTables()
/*
	user := models.CreateUser("1", "1", "a@a")
	fmt.Println(user)
	user.Username = "cambio de nombre 222"
	user.Password = "cambia de pass"
	user.Email = "cambio Email"
	user.Save()*/
	// user.Delete()
	/*
	user := models.GetUsers()
	fmt.Println(user)
	json.Marshal(user)
	models.Closeconnection()
*/

	os.Setenv("HOST", "localhost")
	// eliminar variable de entorno
	//os.Unsetenv("HOST")

	// obtener la variable de entorno
	env := os.Getenv("HOST")
	fmt.Println(env)

	url := config.GetUrlDatabase()
	fmt.Println(url)
}
