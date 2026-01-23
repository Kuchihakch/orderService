package models

type OrderRepository interface {
	Create(order Order) error
	GetByID(id string) (Order, error)
	GetAll() []Order
	GetByUserID(userID string) []Order
	UpdateStatus(id string, status string) error
	Delete(id string) error
}
