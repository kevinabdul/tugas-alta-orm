package routes

import(
	user "ormalta/problem-3.2/controllers/user"

	"ormalta/problem-3.2/middlewares"
)

func registerUserRoutes() {
	e.GET("/users", user.GetUsersController, middlewares.AuthenticateUser)

	e.POST("/users", user.AddUserController)

	e.POST("/login", user.LoginUserController)

	e.GET("/users/:id", user.GetUserByIdController, middlewares.AuthenticateUser, middlewares.CheckId)

	e.PUT("/users/:id", user.EditUserController, middlewares.AuthenticateUser, middlewares.CheckId)

	e.DELETE("/users/:id", user.DeleteUserController, middlewares.AuthenticateUser, middlewares.CheckId)
	
}

