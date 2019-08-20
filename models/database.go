package models
import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Initialize(connectionString string) {
  var err error
  DB, err = gorm.Open("postgres", connectionString)
  if err != nil {
    panic(err.Error())
  }
  if err = DB.DB().Ping(); err != nil {
    panic(err.Error())
  }
  fmt.Println("Connected to Database.")
  //defer DB.Close()
}
