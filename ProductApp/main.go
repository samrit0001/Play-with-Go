package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
}

var products = []Product{
	{ID: 1, Name: "Coffee", Cost: 50},
	{ID: 2, Name: "Apple", Cost: 10},
}

func getproduct(ctx *gin.Context) {
	ctx.JSON(200, products)
}

func createProduct(ctx *gin.Context) {
	var product Product

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid JSON"})
	}

	products = append(products, product)

	ctx.JSON(http.StatusCreated, gin.H{
		"message":         "Product created",
		"Product Details": product,
	})
}

func getProductById(ctx *gin.Context) {
	idparam := ctx.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
		return
	}

	for _, product := range products {
		if product.ID == id {
			ctx.JSON(http.StatusOK, gin.H{
				"message ": "Product found",
				"Product ": product,
			})
			return
		}
	}
	ctx.JSON(200, gin.H{
		"message": "sorry product  not found",
	})
}

func deleteproduct(ctx *gin.Context) {
	idparam := ctx.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			ctx.JSON(200, gin.H{
				"message": "product deleted",
				"Product": product,
			})
			return
		}
	}
	ctx.JSON(200, gin.H{"message": "no product found"})

}

func authmiddle(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token != "motti" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "wrong api key"})
		return
	}
	ctx.Next()
}

func main() {
	//create router
	r := gin.Default()
	//create middlewares
	r.Use(authmiddle)
	r.Use(logmiddle)

	//we create apis

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Product App"})
	})
	r.GET("/getproduct", getproduct)
	r.POST("/create", createProduct)
	r.GET("/getproduct:id", getProductById)
	r.DELETE("/deleteproduct:id", deleteproduct)

	//start the server
	fmt.Println("Server started")
	r.Run(":8080")

}
