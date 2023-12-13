package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type Products struct {
	products []Product
}

func openFile(path string) []Product {
	data, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	dataRead, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}

	slice := []Product{}
	json.Unmarshal(dataRead, &slice)
	return slice
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getProducts() []Product {
	data := openFile("products.json")
	return data
}

func getByID(id string) (Product, error) {
	data := openFile("products.json")
	for _, product := range data {
		if fmt.Sprint(product.Id) == id {
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("product with id '%s' not found", id)
}

func searchByPrice(price string) []Product {
	data := openFile("products.json")
	var products []Product
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Fatal(err)
	}
	for _, product := range data {
		if product.Price >= priceFloat {
			products = append(products, product)
		}
	}
	return products
}

func main() {
	server := gin.Default()

	server.GET("/ping", ping)

	groupProduct := server.Group("/products")

	groupProduct.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, getProducts())
	})

	groupProduct.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		product, err := getByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, product)
	})

	groupProduct.GET("/search/:priceGt", func(ctx *gin.Context) {
		price := ctx.Param("priceGt")
		product := searchByPrice(price)
		ctx.JSON(http.StatusOK, product)
	})

	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
