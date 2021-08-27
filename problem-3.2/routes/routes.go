package routes

import (
	user "ormalta/problem-3.2/controllers/user"
	book "ormalta/problem-3.2/controllers/book"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", user.GetUsersController)

	e.GET("/users/:id", user.GetUserByIdController)

	e.POST("/users", user.AddUserController)

	e.PUT("/users/:id", user.EditUserController)

	e.DELETE("/users/:id", user.DeleteUserController)



	e.GET("/books", user.GetBooksController)

	e.GET("/books/:id", user.GetBookByIdController)

	e.POST("/books", user.AddBookController)

	e.PUT("/books/:id", user.EditBookController)

	e.DELETE("/books/:id", user.DeleteBookController)

	return e
}