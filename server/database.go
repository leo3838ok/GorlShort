package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:password@(mysql)/shortener?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
