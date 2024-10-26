package http

import (
	"net/http"
	"strconv"

	"github.com/Rawipass/product-service/internal/usecase"
	"github.com/Rawipass/product-service/models"
	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	gender := c.Query("gender")
	style := c.Query("style")
	size := c.Query("size")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	perPage, err := strconv.Atoi(c.Query("per_page"))
	if err != nil || perPage <= 0 {
		perPage = 10
	}
	uc := usecase.NewProductUseCase()
	products, err := uc.ListProducts(gender, style, size, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func CreateOrder(c *gin.Context) {
	var orderRequest models.CreateOrderRequest

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uc := usecase.NewProductUseCase()
	orderID, err := uc.CreateOrder(orderRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order_id": orderID})
}

func ListOrders(c *gin.Context) {
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")
	status := c.Query("status")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	perPage, err := strconv.Atoi(c.Query("per_page"))
	if err != nil || perPage <= 0 {
		perPage = 10
	}

	uc := usecase.NewProductUseCase()
	orders, err := uc.ListOrders(start_date, end_date, status, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
