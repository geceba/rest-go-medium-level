package models

import "errors"

type User struct {
  Id int `json:"id"`
  Username string `json:"username"`
  Password string `json:"password"`
}

type Users []User
var users = make(map[int]User)

const userSchema string = `CREATE TABLE users(
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(30) NOT NULL UNIQUE,
        password VARCHAR(60) NOT NULL,
        email VARCHAR(40),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

func SetDefaultUser() {
  user := User{Id: 1, Username: "geceba", Password: "123"}
  users[user.Id] = user
}

func GetUsers() Users {
  list := Users{}

  for _, user := range users {
    list = append(list, user)
  }
  return list
}

func GetUser(userId int) (User, error) {
  if user, ok := users[userId]; ok {
    return user, nil
  }
  return User{}, errors.New("El usuario no se encuentra dentro del mapa")
}

func SaveUser(user User) User {
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

func DeleteUser(id int) {
  delete(users, id)
}
