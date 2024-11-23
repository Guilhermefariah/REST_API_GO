package controller 

import (
	"go-api/model"
	"net/http"
	"github.com/gin-gonic/gin"
)
type ProductController struct {
	//Usecase
}

func NewProductController() ProductController {
	return ProductController{}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID: 1,
			Name: "Coca-Cola",
			Price: 20,
		}
	}
	ctx.JSON(http.StatusOK, products)
}