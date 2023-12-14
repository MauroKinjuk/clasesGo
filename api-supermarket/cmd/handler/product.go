package handler

import (
	"api-supermarket/internal/domain"
	"api-supermarket/internal/product"
	"api-supermarket/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductRouter struct {
	productGroup *gin.RouterGroup
	service      product.ProductService
}

// NewProductRouter returns a new product router
func NewProductRouter(g *gin.RouterGroup) ProductRouter {
	// Slice fills the database with the data from the json file
	slice := pkg.FillDb("../products.json")
	// Create a new product repository
	repo := product.NewProductRepository(slice)
	// Create a new product service
	serv := product.NewProductService(repo)

	return ProductRouter{g, serv}
}

// Routes defines the routes for the product router
func (r *ProductRouter) Routes() {
	r.productGroup.GET("/ping", r.Ping())
	r.productGroup.GET("/", r.GetAll())
	r.productGroup.GET("/:id", r.GetById())
	r.productGroup.POST("/", r.Create())
}

// Ping is a simple ping/pong endpoint
func (r *ProductRouter) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

// GetAll returns all the products
func (r *ProductRouter) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := r.service.GetAll()
		ctx.JSON(200, products)
	}
}

// GetById returns a product by id
func (r *ProductRouter) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "id must be a number",
			})
			return
		}
		data := r.service.GetById(id)
		ctx.JSON(200, data)
	}
}

func (r *ProductRouter) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind the body to the ProductCreate struct
		// If the body is invalid, return a 400
		var body domain.ProductCreate
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid body",
			})
			return
		}

		createdProduct, err := r.service.Create(body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error creating product",
			})
			return
		}
		ctx.JSON(http.StatusCreated, createdProduct)
	}
}
