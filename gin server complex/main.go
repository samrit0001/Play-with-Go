package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Amrit", Email: "abc@g.com"},
	{ID: 2, Name: "motti", Email: "def@g.com"},
}

func getusers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}

func createuser(ctx *gin.Context) {

	var user User
	//Parsing the JSON from request to user variable
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}
	//append the user information
	users = append(users, user)

	//send response for created user
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user is created",
		"user":    user,
	})
}

func getbyid(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	for _, user := range users {
		if user.ID == id {
			ctx.JSON(http.StatusOK, user)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
}

func deleteuser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "User Deleted "})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "user not Found"})
}

func authmiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token != "mysecert123" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong api key"})
		return
	}
	ctx.Next()
}

func loggingmiddleware(ctx *gin.Context) {
	startTime := time.Now()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path

	//Process the Request
	ctx.Next() // pass control to next handler

	//Log Details
	statusCode := ctx.Writer.Status()
	responseTime := time.Since(startTime)

	//Print it
	fmt.Printf("[%s] %s %s - %d (%s)\n",
		time.Now().Format("2006-01-02 15:04:05"), // Timestamp
		method,                                   // HTTP method
		path,                                     // URL path
		statusCode,                               // Status code
		responseTime,                             // Response time
	)
}
func main() {
	r := gin.Default()

	//middleware
	r.Use(loggingmiddleware)
	r.Use(authmiddleware)
	//API Routes
	r.GET("/getusers", getusers)
	r.POST("/createuser", createuser)
	r.GET("/getbyid/:id", getbyid)
	r.POST("/deleteuser/:id", deleteuser)

	//Run the server
	fmt.Println("Started the Server")
	r.Run(":8080")
}
