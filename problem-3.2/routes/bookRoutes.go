package routes

import(
	book "ormalta/problem-3.2/controllers/book"

	"ormalta/problem-3.2/middlewares"
)

func registerBookRoutes() {
	e.GET("/books", book.GetBooksController)

	e.GET("/books/:id", book.GetBookByIdController)

	e.POST("/books", book.AddBookController, middlewares.AuthenticateUser)

	e.PUT("/books/:id", book.EditBookController, middlewares.AuthenticateUser)

	e.DELETE("/books/:id", book.DeleteBookController, middlewares.AuthenticateUser)
}