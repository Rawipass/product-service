package routes

import (
	"github.com/Rawipass/product-service/internal/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/products/list", http.ListProducts)
	router.POST("/order/create", http.CreateOrder)
	router.GET("/order/list", http.ListOrders)

	return router
}
