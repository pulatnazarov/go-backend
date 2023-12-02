package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Store struct {
	Dealer     Dealer
	Repository Repository
	Budget     int
	Profit     int
}

func NewStore(repository Repository) Store {
	return Store{
		Repository: repository,
		Budget:     10_000,
		Profit:     0,
	}
}

func (s *Store) StartSell() {
	productSellRequest := getProductInfo()

	product, exists := s.Repository.Search(productSellRequest.Name)
	if !exists {
		fmt.Printf("We do not have %s product\nWe will bring %s in the next time\n", productSellRequest.Name, productSellRequest.Name)

		product, success := s.Dealer.ProvideProduct(productSellRequest.Name, s.Budget, productSellRequest.Quantity)
		if !success {
			fmt.Println("We will buy!!!!")
			return
		}

		s.Repository.Products.AddProduct(product)

		s.Budget -= product.OriginalPrice * product.Quantity

		return
	}

	if product.Quantity < productSellRequest.Quantity {
		fmt.Printf("We do not have enought %s product, left %d\n", productSellRequest.Name, product.Quantity)
		return
	}

	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)

	s.Profit += productSellRequest.Quantity * (product.Price - product.OriginalPrice)

	s.printStats()
}

func (s *Store) printStats() {
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, '\t', 0)

	fmt.Fprintln(w, "Name\tQuantity\tPrice\tOriginal Price\t")

	for _, product := range s.Repository.Products {
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\n", product.Name, product.Quantity, product.Price, product.OriginalPrice)
	}

	fmt.Fprintf(w, "\t\t\t\tBudget: %d\n", s.Budget)
	fmt.Fprintf(w, "\t\t\t\tProfit: %d\n", s.Profit)

	w.Flush()
}
