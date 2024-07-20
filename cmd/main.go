package main

import (
	"go-api/controller"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	//INICIALIZAR O USECASE
	ProductUseCase := usecase.NewProductUseCase()

	//CAMADA DE CONTROLLERS
	//INICIALIZAR O CONTROLLER
	ProductController := controller.NewProductController(ProductUseCase)
	//MAPEAR A ROTA
	server.GET("/products", ProductController.GetProducts)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}
