package customer

import (
	"fmt"
	"go-backend/store/basket"
	"go-backend/store/product"
)

type Person struct {
	Name string
	Age  int
	LastName string
}

type Customer struct {
	Name   string
	Cash   int
	Basket basket.Basket
}

func NewCustomer(name string, cash int, basket basket.Basket) Customer {
	return Customer{
		Name:   name,
		Cash:   cash,
		Basket: basket,
	}
}

func (c *Customer) AddProduct(p product.Product) {
	c.Basket.ProductList.AddProduct(p)
	for _, p := range c.Basket.ProductList {
		c.Basket.Total += p.Price * p.Quantity
	}
}

func GetCustomerInfo() (name string, cash int) {
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your cash: ")
	fmt.Scan(&cash)

	return
}
