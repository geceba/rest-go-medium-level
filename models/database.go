package models

import (
  "fmt"
  "log"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
const username string = "root"
const password string = "zxczxc.123"
const host string = "localhost"
const port int = 3306
const database string = "project_go"

func CreateConnection() {
  if connection, err := sql.Open("mysql", generateURL()); err != nil {
    panic(err)
  } else {
    db = connection
    fmt.Println("Connection success")
  }
}

func CreateTables() {
  createTable("users", userSchema)
}
func createTable(tableName, schema string){
  if !existsTable(tableName){
    Exec(schema)
  }
}

func CloseConnection() {
  db.Close()
}

func generateURL() string {
  return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}

func existsTable(tableName string) bool {
  sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
  rows, _ := Query(sql)

  return rows.Next()
}

func Exec(query string, args ...interface{})(sql.Result, error) {
  result, err := db.Exec(query, args...)

  if err != nil {
    log.Println(err)
  }

  return result, err
}

func Query(query string, args ...interface{})(*sql.Rows, error) {

  rows, err := db.Query(query, args...)
  if err != nil {
    log.Println(err)
  }

  return rows, err
}

func Ping() {
  if err := db.Ping(); err != nil {
    panic(err)
  }
}
