package repo

import (
	"errors"
	"go-products/models"
)

type InMemoryOrder struct {
	data map[string]*models.Order
}

func NewRepository() *InMemoryOrder {
	return &InMemoryOrder{
		data: make(map[string]*models.Order),
	}
}

func (o *InMemoryOrder) Create(order models.Order) error {
	if _, exist := o.data[order.ID]; exist {
		return errors.New("Product ID Already Exist")
	}
	o.data[order.ID] = &order
	return nil
}

func (o *InMemoryOrder) GetByID(id string) (models.Order, error) {
	if d, exist := o.data[id]; exist {
		return *d, nil
	}
	return models.Order{}, errors.New("Order Not Found")
}

func (o *InMemoryOrder) GetAll() []models.Order {
	res := []models.Order{}
	for _, v := range o.data {
		res = append(res, *v)
	}
	return res
}

func (o *InMemoryOrder) GetByUserID(userID string) []models.Order {
	// via loop
	res := []models.Order{}
	for _, v := range o.data {
		if v.UserID == userID {
			res = append(res, *v)
		}
	}
	return res
}

func (o *InMemoryOrder) UpdateStatus(id string, status string) error {
	if d, exist := o.data[id]; exist {
		d.Status = status
		return nil
	}
	return errors.New("Order Not Found")
}

func (o *InMemoryOrder) Delete(id string) error {
	if _, exist := o.data[id]; exist {
		delete(o.data, id)
		return nil
	}
	return errors.New("Order Not Found")
}
