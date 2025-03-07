package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json: "id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "amrit"},
	{ID: 2, Name: "motti"},
}

func getusers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)

}

func createuser(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "INVALID JSON",
		})
		return
	}

	//append the user to in-memory slice
	users = append(users, user)

	//send Response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user":    user,
	})

}

func main() {

	//router
	r := gin.Default() //default router with logging and recovery

	r.GET("/getusers", getusers)
	r.POST("/createuser", createuser)

	fmt.Println("Starting the server")
	r.Run(":8080")
}
