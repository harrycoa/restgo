package models
type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
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
func NewUser(username, password string) *User {
	user := &User {Username: username, Password: password}
	return user
}

func (this *User) Save(){
	sql := "INSERT users SET username=?, password=?"
	Exec(sql, this.Username,this.Password)
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