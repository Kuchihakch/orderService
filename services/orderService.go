package services

import (
	"errors"
	"go-products/models"
	"time"
)

type OrderServiceImpl struct {
	OrderRepo   models.OrderRepository // service depedence ke OrderRepo, productRepo
	ProductRepo models.ProductRepository
}

func (r *OrderServiceImpl) CreateOrder(order models.Order) error {
	if order.ID == "" || order.Amount <= 0 || order.UserID == "" || order.ProductID == "" {
		return errors.New("Invalid Order Payload")
	}
	_, err := r.OrderRepo.GetByID(order.ID)
	if err == nil {
		return errors.New("Order ID Already Exist")
	}

	//cek produk
	_, errProduct := r.ProductRepo.GetByID(order.ProductID)
	if errProduct != nil {
		return errProduct
	}

	//update stock nanti ketika udah PAID + validasi stok

	order.Status = "NEW"
	now := time.Now()
	y, m, d := now.Date()
	order.CreatedAt = time.Date(y, m, d, 0, 0, 0, 0, now.Location())
	return r.OrderRepo.Create(order)
}

func (r *OrderServiceImpl) PayOrder(id string) error {
	if id == "" {
		return errors.New("ID Required")
	}
	d, err := r.OrderRepo.GetByID(id)
	if err != nil {
		return err
	}
	if d.Status == "PAID" {
		return errors.New("Cannot Update Status Successful Transaction")
	}
	p, errProduct := r.ProductRepo.GetByID(d.ProductID)
	if errProduct != nil {
		return errProduct
	}
	if p.Stock < d.Amount {
		return errors.New("Not Enough Product Stock")
	}
	newStock := p.Stock - d.Amount
	errUpdate := r.ProductRepo.UpdateStock(p.ID, newStock)
	if errUpdate != nil {
		return errUpdate
	}
	return r.OrderRepo.UpdateStatus(id, "PAID")
}

func (r *OrderServiceImpl) CancelOrder(id string) error {
	if id == "" {
		return errors.New("ID Required")
	}
	d, err := r.OrderRepo.GetByID(id)
	if err != nil {
		return err
	}
	if d.Status == "PAID" {
		return errors.New("Cannot Update Status Successful Transaction")
	}
	return r.OrderRepo.UpdateStatus(id, "CANCELLED")
}

func (r *OrderServiceImpl) FindAll() []models.Order {
	return r.OrderRepo.GetAll()
}

func (r *OrderServiceImpl) FindById(id string) (models.Order, error) {
	if id == "" {
		return models.Order{}, errors.New("ID Required")
	}
	return r.OrderRepo.GetByID(id)
}

func (r *OrderServiceImpl) FindByUserID(userId string) []models.Order {
	return r.OrderRepo.GetByUserID(userId)
}

func (r *OrderServiceImpl) DeleteOrder(id string) error {
	if id == "" {
		return errors.New("ID Required")
	}
	return r.OrderRepo.Delete(id)
}
