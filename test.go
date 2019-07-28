package main

import(
  "./models"
)

func main() {
  models.CreateConnection()


  models.CreateTables()
  models.CloseConnection()
}
