package product

import (
	"fmt"
	"math/rand"
)

var (
	List = ProductList{
		{
			Name:          "Non",
			Quantity:      10,
			Price:         4_000,
			OriginalPrice: 2_000,
		},
		{
			Name:          "Cola",
			Quantity:      15,
			Price:         13_000,
			OriginalPrice: 6_500,
		},
		{
			Name:          "Nestle",
			Quantity:      20,
			Price:         3_000,
			OriginalPrice: 1_500,
		},
	}
)

type Product struct {
	Name          string
	Quantity      int
	Price         int
	OriginalPrice int
}

type ProductSellRequest struct {
	Name     string
	Quantity int
}

func NewProduct(name string, price, quantity int) Product {
	return Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

type ProductList []Product

func (p *ProductList) AddProduct(product Product) {
	*p = append(*p, product)
}

func (p *ProductList) RemoveProduct(index int) {
	*p = append((*p)[:index], (*p)[index+1:]...)
}

func GetProductInfo() ProductSellRequest {
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

func GenerateProductPrice(min, max int) int {
	return rand.Intn(max-min) + min
}
