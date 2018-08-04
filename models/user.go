package models

import "fmt"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const userSchema string = `create database users(
  id int(6) unsiged auto_increment primary key,
  username varchar(100) not null,
  password varchar(256) not null,
  email varchar(200) null,
  create_date timestamp default current_timestamp
);`

type Users []User

// crear constructos
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

func GetUser(id int) *User{
	user := NewUser("","","")
	sql := "select id, username, password, email FROM users where id=?"
	rows, _ := Query(sql,id)
	for rows.Next(){
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Password)
	}
	return user
}

func GetUsers() Users {
	sql := "Select id,username,password,email from users"
	users := Users{}
	rows, _ := Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Password)
		users = append(users,user)
	}
	return users
}

func (this *User) Save() {
	if this.Id == 0 {
		this.insert()
	} else {
		this.update()
	}
}

func (this *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := Exec(sql, this.Username, this.Password, this.Email)
	this.Id, _ = result.LastInsertId()
}

func (this *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=?"
	Exec(sql, this.Username, this.Password, this.Email)
}

func (this *User) Delete() {
	sql3 := "DELETE FROM users WHERE id=?"
	Exec(sql3, this.Id)
	fmt.Println("se elimino correctamente")
}

/*
var users = make(map[int]User)

func SetDefaultUser(){
	user := User {Id: 1, Username:"harry12", Password:"1234"}
	users[user.Id] = user
}
func GetUsers() Users{
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

func GetUser(userId int)(User, error){
	if user, ok := users[userId]; ok{
		return user, nil
	}
	return User{}, errors.New("el usuario no se encuentra en el mapa.")
}

func SaveUser(user User) User{
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}

func UpdateUser(user User, username, password string) User {
	user.Username = username
	user.Password = password
	users[user.Id] = user
	return user
}

func DeleteUser(id int){
	delete(users,id)
}
*/
