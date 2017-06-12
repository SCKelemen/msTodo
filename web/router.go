package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/jwt"
)

func ApiRouter() *gin.Engine {
	router := gin.Default()
	
	router.StaticFile("/robots.txt", "robots.txt")
	router.StaticFile("/favicon.ico", "favicon.ico")

	api := router.Group("/api/v1/")
	// List
	api.GET("/l", GetLists)			//List Todo Lists
	api.GET("/l/:id", GetList)		//Get Todo List by ID
	api.POST("/l", CreateList)		//Create Todo List
	api.PATCH("/l/:id", UpdateList)		//Update Todo List
	api.DELETE("/l/:id", RemoveList)	//Delete Todo List

	// Todo
	api.GET("/t", GetTodos)			//List All Todos
	api.GET("/t/:id", GetTodo)		//Get Todo by ID
	api.POST("/t", CreateTodo)		//Create Todo
	api.PATCH("/t/:id", UpdateTodo)		//Update Todo
	api.DELETE("/t/:id", RemoveTodo)	//Delete Todo

	// Comment
	// [todo]

	// Reactions
	// [todo]

	

	// To use auth routes you must have had a JWT token issued to you
	// and requests must have the following header:
	// -- Authorization: Bearer `token`
	auth := api.Group("/auth/")
	// Use a JWT middleware authentication handler to only allow requests 
	// that have been signed by our 'mysupersecretpassword' secret
	auth.Use(jwt.Auth("mysupersecretpassword"))
	auth.GET("/resetpassword", GetResetPassword)
	auth.POST("/resetpassword", PostResetPassword)
	auth.POST("/image", PostImage)
	
	return router;
}
