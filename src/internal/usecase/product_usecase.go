package usecase

import (
	"github.com/Rawipass/product-service/internal/repository"
	"github.com/Rawipass/product-service/models"
)

type ProductUsecase struct {
	repo repository.ProductRepository
}

func NewProductUseCase() *ProductUsecase {
	usecase := ProductUsecase{}
	return &usecase
}

func (uc *ProductUsecase) ListProducts(gender, style, size string, page, perPage int) ([]models.Product, error) {
	repo := repository.NewProductRepository()
	products, err := repo.ListProducts(gender, style, size, page, perPage)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (uc *ProductUsecase) CreateOrder(orderRequest models.CreateOrderRequest) (int, error) {
	orderID, err := uc.repo.CreateOrder(orderRequest)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (uc *ProductUsecase) ListOrders(start_date, end_date, status string, page, perPage int) ([]models.Order, error) {
	repo := repository.NewProductRepository()
	orders, err := repo.ListOrders(start_date, end_date, status, page, perPage)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
