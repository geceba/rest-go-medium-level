package config

import (
  "github.com/eduardogpg/gonv"
  "fmt"
)

type DatabaseConfig struct {
  Username string
  Password string
  Host string
  Port int
  Database string
}

var database *DatabaseConfig

func init(){
  database = &DatabaseConfig{}
  database.Username = gonv.GetStringEnv("USERNAME", "root")
  database.Password = gonv.GetStringEnv("PASSWORD", "zxczxc.123")
  database.Host = gonv.GetStringEnv("HOST", "localhost")
  database.Port = gonv.GetIntEnv("PORT", 3306)
  database.Database = gonv.GetStringEnv("DATABASE","project_go")
}

func (this *DatabaseConfig) url() string {
  return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", this.Username, this.Password, this.Host, this.Port, this.Database)
}

func GetUrlDatabase() string{
  return database.url()
}
