package main

import(
  "fmt"
  "./orm"
)

func main() {
  orm.CreateConnection()
  //orm.CreateTables()

  //user := orm.CreateUser("Gerardo", "123", "ger@bio.com")
  //user.Save()
  user := orm.GetUser(1)
  fmt.Println(user)

  orm.CloseConnection()
}
