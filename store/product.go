package main

import "fmt"

var (
	productList = ProductList{
		{
			Name:     "Non",
			Price:    4_000,
			Quantity: 10,
		},
		{
			Name:     "Cola",
			Price:    13_000,
			Quantity: 15,
		},
		{
			Name:     "Nestle",
			Price:    3_000,
			Quantity: 20,
		},
	}
)

type Product struct {
	Name     string
	Price    int
	Quantity int
}

type ProductSellRequest struct {
	Name     string
	Quantity int
}

type ProductList []Product

func NewProduct(name string, price, quantity int) Product {
	return Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

func getProductInfo() ProductSellRequest {
	var (
		productName string
		quantity    = 0
	)

	fmt.Print("Enter product name: ")
	fmt.Scan(&productName)

	fmt.Print("Enter product quantity: ")
	fmt.Scan(&quantity)
	fmt.Println()

	return ProductSellRequest{
		Name:     productName,
		Quantity: quantity,
	}
}
