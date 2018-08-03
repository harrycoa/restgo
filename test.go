package main
import (
	"./models"
)
func main(){
	models.CreateConnection()
	 models.Ping()
	//result := models.ExistTable("users")
	//fmt.Println(result)
	//models.CreateTables()

	/*
	user := models.NewUser("aaa", "123")
	fmt.Println(user)
	user.Save()

	*/
	models.Closeconnection()
}
