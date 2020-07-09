package server

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MySQL *gorm.DB

func InitMySQL() {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s",
		os.Getenv("MYSQL_USER_NAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	MySQL, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
