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



	e.GET("/books", book.GetBooksController)

	e.GET("/books/:id", book.GetBookByIdController)

	e.POST("/books", book.AddBookController)

	e.PUT("/books/:id", book.EditBookController)

	e.DELETE("/books/:id", book.DeleteBookController)

	return e
}