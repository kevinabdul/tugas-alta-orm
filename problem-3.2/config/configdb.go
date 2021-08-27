package config

import (
	"fmt"
	"ormalta/problem-3.2/models"
	
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	Db *gorm.DB
)

func InitDb() {
	config := map[string]interface{}{
		"DB_USERNAME": "untukalta",
		"DB_PASSWORD": "tugasorm",
		"DB_DATABASE": "user",
		"DB_PORT": 3306,
		"DB_HOST": "localhost",
		"DB_NAME": "tugasorm"}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"])

	var err error
	Db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Book{})

}