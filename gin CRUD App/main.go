package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var products []Product
var idcounter = 1
var mu sync.Mutex

func getproducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, products)
}

func createProduct(ctx *gin.Context) {
	var newproduct Product
	mu.Lock()
	defer mu.Unlock()
	ctx.ShouldBindJSON(&newproduct)

	newproduct.ID = idcounter
	idcounter++
	products = append(products, newproduct)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product Created",
		"Product": newproduct,
	})
}

func getProductByID(ctx *gin.Context) {
	idfetch := ctx.Param("id")
	id, err := strconv.Atoi(idfetch)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
	}

	for _, product := range products {
		if product.ID == id {
			ctx.JSON(http.StatusOK, product)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
}

func updateProduct(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	idfetch := ctx.Param("id")
	id, err := strconv.Atoi(idfetch)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
	}

	var updatedProduct Product
	ctx.ShouldBindJSON(&updatedProduct)

	for i, product := range products {
		if product.ID == id {
			products[i].Name = updatedProduct.Name
			products[i].Quantity = updatedProduct.Quantity
			ctx.JSON(http.StatusOK, products[i])
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
}

func deleteProduct(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	idfetch := ctx.Param("id")
	id, err := strconv.Atoi(idfetch)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
}

func workers(products []Product, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, product := range products {
		sum = sum + product.Quantity
	}
	results <- sum
}

func calculatesum(products []Product, numworkers int) int {
	//create a channel to receive results
	resultchan := make(chan int, numworkers)

	// Create Partition sizes for each worker
	blocksize := (len(products) + numworkers - 1) / numworkers

	// create waitgroups to track our workers
	var wg sync.WaitGroup

	//start the workers and provide the details of blocksize
	for i := 0; i < numworkers; i++ {
		start := i * blocksize
		end := start + blocksize
		if end > len(products) {
			end = len(products)
		}
		if start > len(products) {
			break
		}

		wg.Add(1)
		go workers(products[start:end], resultchan, &wg)
	}

	//close our channels
	go func() {
		wg.Wait()
		close(resultchan)
	}()

	total := 0
	for sum := range resultchan {
		total = total + sum
	}
	return total

}

func getquantity(ctx *gin.Context) {

	//create workers depending on my CPU cores
	numworkers := 3

	finalquantity := calculatesum(products, numworkers)

	ctx.JSON(http.StatusOK, gin.H{"Quantity": finalquantity})

}

func main() {

	r := gin.Default()

	r.GET("/products", getproducts)
	r.GET("/products/:id", getProductByID)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)
	r.GET("/quantity", getquantity)

	fmt.Println("Server Started")

	r.Run(":8080")
}
