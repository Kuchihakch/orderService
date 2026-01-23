package main

import (
	"fmt"
	"go-products/models"
	"go-products/repo"
	"go-products/services"
)

func main() {
	repo := repo.NewRepository()
	services := services.OrderServiceImpl{Repo: repo}

	order1 := models.Order{
		ID:     "order-1",
		Amount: 3,
		UserID: "user-1",
	}
	services.CreateOrder(order1)
	fmt.Println(services.FindAll())
	fmt.Println(services.FindByUserID("user-1"))
	fmt.Println(services.FindById("order-2")) // notfound
	services.DeleteOrder("order-1")
	fmt.Println(services.FindAll())
	fmt.Println(services.FindById("order-1"))
}
