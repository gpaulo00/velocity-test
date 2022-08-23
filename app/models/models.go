package models

import (
	"fmt"

	"github.com/gpaulo00/velocity-test/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func Init() {
  config    := config.GetConfig()
  if config.GetString("server.database") == "mysql" {
    // Declare variables configuration mysql
    host_mysql      := config.GetString("mysql.host")
    port_mysql      := config.GetString("mysql.port")
    database_mysql  := config.GetString("mysql.database")
    username_mysql  := config.GetString("mysql.username")
    password_mysql  := config.GetString("mysql.password")
    db, err = gorm.Open("mysql", username_mysql + ":" + password_mysql +"@tcp(" + host_mysql + ":" + port_mysql + ")/" + database_mysql + "?charset=utf8&parseTime=True&loc=Local")
  }

  if err != nil {
    fmt.Println(err)
    fmt.Println(db)
  }
}

func GetDB() *gorm.DB {
	return db
}
