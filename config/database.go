package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBMASTER *gorm.DB

func ConnectDb() {
	database, err := gorm.Open(mysql.Open("adextux:adexganteng@tcp(localhost:3306)/pos_member"))
	if err != nil {
		panic(err)
	}
	DB = database
}

func ConnectDbmaster() {
	dbmaster, err := gorm.Open(mysql.Open("adextux:adexganteng@tcp(localhost:3306)/pos_master"))
	if err != nil {
		panic(err)
	}
	DBMASTER = dbmaster
}
