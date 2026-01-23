package models

type OrderService interface {
	CreateOrder(order Order) error
	PayOrder(id string) error
	CancelOrder(id string) error
	GetOrdersByUser(userID string) ([]Order, error)
}
