package main

import(
  _ "fmt"
  _ "./config"
  "./models"
)

func main() {
  models.CreateConnection()
  models.Ping()
}
