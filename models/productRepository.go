package models

type ProductRepository interface {
	Create(product Product) error
	GetByID(id string) (Product, error)
	UpdateStock(id string, stock int) error
}
