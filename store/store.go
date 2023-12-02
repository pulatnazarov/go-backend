package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Store struct {
	Repository Repository
	Cash       int
}

func NewStore(repository Repository) Store {
	return Store{
		Repository: repository,
		Cash:       0,
	}
}

func (s *Store) StartSell() {
	productSellRequest := getProductInfo()

	product, exists := s.Repository.Search(productSellRequest.Name)
	if !exists {
		fmt.Printf("We do not have %s product\n", productSellRequest.Name)
		return
	}

	if product.Quantity < productSellRequest.Quantity {
		fmt.Printf("We do not have enought %s product, left %d\n", productSellRequest.Name, product.Quantity)
		return
	}

	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)

	s.Cash += productSellRequest.Quantity * product.Price
}

func (s *Store) printStats() {
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, '\t', 0)
	fmt.Fprintln(w, "Name\tQuantity\tPrice\t")
	for _, product := range s.Repository.Products {
		fmt.Fprintf(w, "%s\t%d\t%d\n", product.Name, product.Quantity, product.Price)
	}
	fmt.Fprintf(w, "\t\t\t\tCash: %d\n", s.Cash)
	w.Flush()
}
