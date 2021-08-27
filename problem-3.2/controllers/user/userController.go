package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	user "ormalta/problem-3.2/services/user"
	models "ormalta/problem-3.2/models"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	fmt.Println("Inside getusersController")
	users, err := user.GetUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserByIdController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))

	targetUser, rowsAffected, err := user.GetUserById(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	return c.JSON(http.StatusOK, targetUser)
}

func AddUserController(c echo.Context) error {
	newUser := models.User{}
	c.Bind(&newUser)

	res, err := user.AddUser(&newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.User
	}{Status: "succes", Message: "User has been created!", User: res})

}

func EditUserController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	newData := models.User{}
	c.Bind(&newData)

	edittedUser, rowsAffected, err := user.EditUser(newData, targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.User
	}{Status: "succes", Message: "User has been updated!", User: edittedUser})
}

func DeleteUserController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	deleted, rowsAffected, err := user.DeleteUser(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.User
	}{Status: "succes", Message: "User has been deleted!", User: deleted})

}
