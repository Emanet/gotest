package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB {
	user := "root"
	pass := "password"
	host := "localhost"
	dbName := "todo"

	credentials := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbName)

	db, err := gorm.Open("mysql", credentials)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
