package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int
	Name string
}

var users = []User{
	{ID: 1, Name: "ammy"},
}

func getusers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "getusers page"})
}

func main() {
	//creae default router
	r := gin.Default()

	//gin.Default is equal to the below one
	// r := gin.New()
	// r.Use(gin.Logger(), gin.Recovery())

	//create api route

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{ //gin.H{} is a shorthand for map[string]interface{}
			"message": "hey this is main page",
			"status":  "page working",
		})
	})
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, users)
	})

	r.GET("/users1", getusers)

	r.Run(":8080")
}
