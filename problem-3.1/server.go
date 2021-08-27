package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type User struct {
	gorm.Model
	Id     		int 	`gorm:"primaryKey`
	Name   		string	`json:name form:name`
	Email 		string	`gorm:"unique" json:email form:email`
	Password 	string	`json:password form: password`
}

var (
	db *gorm.DB
)

func initDb() {
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
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

}

func main() {
	initDb()

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/users", getusers)

	e.GET("/users/:id", getUserById)

	e.POST("/users", addUser)

	e.PUT("/users/:id", editUser)

	e.DELETE("/users/:id", deleteUser)

	e.Start(":8000")
}

func getusers(c echo.Context) error {
	var users []User
	res := db.Find(&users)

	if res.Error != nil {
		return c.JSON(http.StatusBadRequest, 123)
	}
	return c.JSON(http.StatusOK, users)
}

func getUserById(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	var user User

	res := db.Find(&user, targetId)

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	return c.JSON(http.StatusOK, user)
}

func addUser(c echo.Context) error {
	newUser := User{}
	c.Bind(&newUser)

	db.Create(&newUser)
	
	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User User
	}{Status: "succes", Message: "User has been created!", User: newUser})

}

func editUser(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	targetUser := &User{}
	res := db.Find(targetUser, targetId)

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}

	c.Bind(targetUser)

	res = db.Updates(targetUser)

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User User
	}{Status: "succes", Message: "User has been updated!", User: *targetUser})
}

func deleteUser(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	targetUser := &User{}
	res := db.Find(targetUser, targetId)

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	
	deleted := *targetUser
	db.Delete(targetUser, targetId)

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User User
	}{Status: "succes", Message: "User has been deleted!", User: deleted})

}
