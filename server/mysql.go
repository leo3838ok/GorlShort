package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MySQL *gorm.DB

func InitMySQL() {
	var err error
	MySQL, err = gorm.Open("mysql", "root:password@(mysql)/shortener?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
