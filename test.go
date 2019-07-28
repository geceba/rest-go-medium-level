package main

import(
  "fmt"
  "./models"
)

func main() {
  models.CreateConnection()


  models.CreateTables()

  models.CreateUser("adriel", "123", "hola@hola.com")
  models.CreateUser("juan", "123", "hola@hola.com")
  models.CreateUser("test", "123", "hola@hola.com")

  user := models.GetUsers()
  fmt.Println(user)
  //user.Username = "gerardo"
  //user.Save()
  //user.Delete()
  models.CloseConnection()
}
