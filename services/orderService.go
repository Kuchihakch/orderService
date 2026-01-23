package services

import (
	"errors"
	"go-products/models"
	"time"
)

type OrderServiceImpl struct {
	Repo models.OrderRepository // service depedence ke repo
}

func (r *OrderServiceImpl) CreateOrder(order models.Order) error {
	if order.ID == "" || order.Amount <= 0 || order.UserID == "" {
		return errors.New("Invalid Order Payload")
	}
	_, err := r.Repo.GetByID(order.ID)
	if err == nil {
		return errors.New("Order ID Already Exist")
	}
	order.Status = "NEW"
	now := time.Now()
	y, m, d := now.Date()
	order.CreatedAt = time.Date(y, m, d, 0, 0, 0, 0, now.Location())
	return r.Repo.Create(order)
}

func (r *OrderServiceImpl) PayOrder(id string) error {
	if id == "" {
		return errors.New("ID Required")
	}
	d, err := r.Repo.GetByID(id)
	if err != nil {
		return err
	}
	if d.Status == "PAID" {
		return errors.New("Cannot Update Status Successful Transaction")
	}
	return r.Repo.UpdateStatus(id, "PAID")
}

func (r *OrderServiceImpl) CancelOrder(id string) error {
	if id == "" {
		return errors.New("ID Required")
	}
	d, err := r.Repo.GetByID(id)
	if err != nil {
		return err
	}
	if d.Status == "PAID" {
		return errors.New("Cannot Update Status Successful Transaction")
	}
	return r.Repo.UpdateStatus(id, "CANCELLED")
}

func (r *OrderServiceImpl) FindAll() []models.Order {
	return r.Repo.GetAll()
}

func (r *OrderServiceImpl) FindById(id string) (models.Order, error) {
	if id == "" {
		return models.Order{}, errors.New("ID Required")
	}
	return r.Repo.GetByID(id)
}

func (r *OrderServiceImpl) FindByUserID(userId string) []models.Order {
	return r.Repo.GetByUserID(userId)
}

func (r *OrderServiceImpl) DeleteOrder(id string) error {
	if id == "" {
		return errors.New("ID Required")
	}
	return r.Repo.Delete(id)
}
