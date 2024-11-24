package main

import (
	"go-api/controller"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductUseCase := usecase.NewProductUseCase()
	//Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
