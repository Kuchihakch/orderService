package services

import (
	"errors"
	"go-products/models"
)

type ProductServiceImpl struct {
	ProductRepo models.ProductRepository
}

// beberapa service yang mungkin akan dipakai admin

func (s *ProductServiceImpl) CreateProduct(product models.Product) error {
	if product.ID == "" || product.Name == "" || product.Stock < 0 {
		return errors.New("Invalid Create Product Payload")
	}
	_, err := s.ProductRepo.GetByID(product.ID)
	if err == nil {
		return errors.New("Product ID Already Exist")
	}
	return s.ProductRepo.Create(product)
}

func (s *ProductServiceImpl) GetProductByID(id string) (models.Product, error) {
	if id == "" {
		return models.Product{}, errors.New("ID Required")
	}
	// return data, err
	return s.ProductRepo.GetByID(id)
}

func (s *ProductServiceImpl) UpdateProductStock(id string, newStock int) error {
	if id == "" || newStock <= 0 {
		return errors.New("Invalid Update Product Stock Payload")
	}
	return s.ProductRepo.UpdateStock(id, newStock)
}
