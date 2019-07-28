package orm

import (
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
  "../config"
)

var db *gorm.DB

func CreateConnection() {
  url := config.GetUrlDatabase()
  if connection, err := gorm.Open("mysql", url); err != nil {
    panic(err)
  } else {
    db = connection
  }
}


func CloseConnection() {
  db.Close()
}

func CreateTables(){
  //db.DropTableIfExists(&User{})
  db.CreateTable(&User{})
}
