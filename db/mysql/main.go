package main

import (
	"log"

	"github.com/jinzhu/gorm"
	gorm_ "github.com/xm-tech/go-lan-demo/db/mysql/gorm_"
)

type User struct {
	Id     int64
	UserId int64
	Name   string
	Sex    int
}

func main() {
	var db *gorm.DB = gorm_.GetDb()

	if !db.HasTable(&User{}) {
		log.Println("create table user")
		db.AutoMigrate(&User{})
		db.Model(&User{}).AddIndex("i_user_id", "user_id")
	}

	user0 := User{UserId: 1, Name: "maxm", Sex: 1}
	user1 := User{UserId: 2, Name: "lfj", Sex: 0}
	db.Create(&user0)
	db.Create(&user1)
}
