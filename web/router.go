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
	api.GET("/l", ListTodo)                 //LIST List

        api.POST("/l", CreateList)              //CREATE List
        api.GET("/l/:id", RetrieveList)         //RETRIEVE List 
        api.PATCH("/l/:id", AlterList)          //ALTER List 
        api.DELETE("/l/:id", PurgeList)         //PURGE List 


	// Todo
	api.GET("/t", ListTodo)			//LIST Todos

	api.POST("/t", CreateTodo)              //CREATE Todo
        api.GET("/t/:id", RetrieveTodo)         //RETRIEVE Todo
        api.PATCH("/t/:id", AlterTodo)          //ALTER Todo
        api.DELETE("/t/:id", PurgeTodo)         //PURGE Todo

	// Labels
	api.GET("/z", ListLabel)		//LIST Labels

	api.POST("/z", CreateLabel)		//CREATE Label
	api.GET("/z/:id", RetrieveLabel)	//RETRIEVE Label
	api.PATCH("/z/:id", AlterLabel)		//ALTER	Label
	api.DELETE("/z/:id", PurgeLabel)	//PURGE	Label

	// Categories
	// [todo]

	// Statuses
	// [todo]

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
