package main

import "fmt"

type Store struct {
	Repository Repository
	Cash       int
}

func (s Store) NewStore(repository Repository) Store {
	return Store{
		Repository: repository,
		Cash:       0,
	}
}

func (s Store) StartSell() {
	s.printStats()

	productSellRequest := getProductInfo()

	product, exists := s.Repository.Search(productSellRequest.Name)
	if !exists {
		fmt.Printf("We do not have %s product", productSellRequest.Name)
		return
	}

	if product.Quantity < productSellRequest.Quantity {
		fmt.Printf("We do not have enought %s product, left %d", productSellRequest.Name, product.Quantity)
		return
	}

	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)

	s.Cash += productSellRequest.Quantity * product.Price

	s.printStats()
}

func (s Store) printStats() {
	fmt.Println("-----------------------------------")
	fmt.Println("CASH: ", s.Cash)
	for _, product := range s.Repository.Products {
		fmt.Println(product.Name, product.Quantity, product.Price)
		fmt.Println()
	}
	fmt.Println("-----------------------------------")
}
