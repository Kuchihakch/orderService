package repo

import (
	"errors"
	"go-products/models"
)

type InMemoryProduct struct {
	data map[string]*models.Product
}

func NewProductRepository() *InMemoryProduct {
	return &InMemoryProduct{
		data: make(map[string]*models.Product),
	}
}

func (p *InMemoryProduct) Create(product models.Product) error {
	// tetap cek keberadaan, karena nanti implementasinya entah di luar product service
	if _, exist := p.data[product.ID]; exist {
		return errors.New("Product ID Already Exist")
	}
	p.data[product.ID] = &product
	return nil
}

func (p *InMemoryProduct) GetByID(id string) (models.Product, error) {
	if d, exist := p.data[id]; exist {
		return *d, nil
	}
	return models.Product{}, errors.New("Product Not Found")
}

func (p *InMemoryProduct) UpdateStock(id string, stock int) error {
	if d, exist := p.data[id]; exist {
		d.Stock = stock
		return nil
	}
	return errors.New("Product Not Found")
}
