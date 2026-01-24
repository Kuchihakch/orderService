package main

import (
	"fmt"
	"go-products/models"
	"go-products/repo"
	"go-products/services"
)

func main() {
	productRepo := repo.NewProductRepository()
	productService := services.ProductServiceImpl{ProductRepo: productRepo}

	orderRepo := repo.NewRepository()
	orderService := services.OrderServiceImpl{OrderRepo: orderRepo, ProductRepo: productRepo}

	// bikin product dulu (as Admin)
	product1 := models.Product{
		ID:    "product-1",
		Name:  "komik my kisah",
		Stock: 3,
	}
	product2 := models.Product{
		ID:    "product-2",
		Name:  "tips n trik kripto",
		Stock: 1,
	}
	productService.CreateProduct(product1)
	productService.CreateProduct(product2)

	fmt.Println(productService.GetProductByID("product-1"))

	//order as User
	order1 := models.Order{
		ID:        "order-1",
		ProductID: "product-1",
		UserID:    "user-1",
		Amount:    2,
	}
	// melebihi stock product 2
	order2 := models.Order{
		ID:        "order-2",
		ProductID: "product-2",
		UserID:    "user-2",
		Amount:    2,
	}
	// aman race, update stock ketika pay
	order3 := models.Order{
		ID:        "order-3",
		ProductID: "product-1",
		UserID:    "user-2",
		Amount:    2,
	}
	// notfound product
	order4 := models.Order{
		ID:        "order-4",
		ProductID: "product-5",
		UserID:    "user-1",
		Amount:    2,
	}

	orderService.CreateOrder(order1)
	orderService.CreateOrder(order2)
	orderService.CreateOrder(order3)
	orderService.CreateOrder(order4)

	fmt.Println(orderService.FindAll())

	fmt.Println(orderService.PayOrder("order-1"))
	//cek race order yang sama
	fmt.Println(orderService.PayOrder("order-3"))

	//admin update dulu stok
	productService.UpdateProductStock("product-1", 5)
	fmt.Println(orderService.PayOrder("order-3"))
}
