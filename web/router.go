package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/mattn/go-colorable"

)

func ApiRouter() *gin.Engine {
	router := gin.Default()
	gin.DefaultWriter = colorable.NewColorableStdout()
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

	// Comment
        api.GET("/x", ListComment)                //LIST Comments

        api.POST("/x", CreateComment)             //CREATE Comment
        api.GET("/x/:id", RetrieveComment)        //RETRIEVE Comment
        api.PATCH("/x/:id", AlterComment)         //ALTER Comment
        api.DELETE("/x/:id", PurgeComment)        //PURGE Comment

	// Reaction
        api.GET("/r", ListReaction)                //LIST Reactions

        api.POST("/r", CreateReaction)             //CREATE Reaction
        api.GET("/r/:id", RetrieveReaction)        //RETRIEVE Reaction
        api.PATCH("/r/:id", AlterReaction)         //ALTER Reaction
        api.DELETE("/r/:id", PurgeReaction)        //PURGE Reaction

	// Statuses
        api.GET("/s", ListStatus)                //LIST Statuses

        api.POST("/s", CreateStatus)             //CREATE Status
        api.GET("/s/:id", RetrieveStatus)        //RETRIEVE Status
        api.PATCH("/s/:id", AlterStatus)         //ALTER Status
        api.DELETE("/s/:id", PurgeStatus)        //PURGE Status

	// Categories
        api.GET("/c", ListCategory)                //LIST Categories

        api.POST("/c", CreateCategory)             //CREATE Category
        api.GET("/c/:id", RetrieveCategory)        //RETRIEVE Category
        api.PATCH("/c/:id", AlterCategory)         //ALTER Category
        api.DELETE("/c/:id", PurgeCategory)        //PURGE Category

	// Users
        api.GET("/u", ListUser)                //LIST Users

        api.POST("/u", CreateUser)             //CREATE User
        api.GET("/u/:id", RetrieveUser)        //RETRIEVE User
        api.PATCH("/u/:id", AlterUser)         //ALTER User
        api.DELETE("/u/:id", PurgeUser)        //PURGE User
	


	// To use auth routes you must have had a JWT token issued to you
	// and requests must have the following header:
	// -- Authorization: Bearer `token`
	auth := api.Group("/auth/")
	// Use a JWT middleware authentication handler to only allow requests 
	// that have been signed by our 'mysupersecretpassword' secret
	auth.Use(jwt.Auth("mysupersecretpassword"))
	//auth.GET("/resetpassword", GetResetPassword)
	//auth.POST("/resetpassword", PostResetPassword)
	//auth.POST("/image", PostImage)
	
	return router;
}
