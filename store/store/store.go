package store

import (
	"fmt"
	"go-backend/store/customer"
	"go-backend/store/dealer"
	"go-backend/store/product"
	"go-backend/store/repository"
	"os"
	"text/tabwriter"
)

type Store struct {
	Dealer     dealer.Dealer
	Repository repository.Repository
	Budget     int
	Profit     int
}

func NewStore(repository repository.Repository) Store {
	return Store{
		Repository: repository,
		Budget:     10_000,
		Profit:     0,
	}
}

func (s *Store) StartSell(user customer.Customer) {
	productSellRequest := product.GetProductInfo()

	p, exists := s.Repository.Search(productSellRequest.Name)
	if !exists {
		fmt.Printf("We do not have %s p\nWe will bring %s in the next time\n", productSellRequest.Name, productSellRequest.Name)

		product, success := s.Dealer.ProvideProduct(productSellRequest.Name, s.Budget, productSellRequest.Quantity)
		if !success {
			fmt.Println("We will buy!!!!")
			return
		}

		s.Repository.Products.AddProduct(product)

		s.Budget -= product.OriginalPrice * product.Quantity

		return
	}

	if user.Cash < p.Price*productSellRequest.Quantity {
		fmt.Println("You don't have enough cash")
		return
	}

	if p.Quantity < productSellRequest.Quantity {
		fmt.Printf("We do not have enought %s p, left %d\n", productSellRequest.Name, p.Quantity)
		return
	}

	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)

	user.AddProduct(product.Product{
		Name:     p.Name,
		Quantity: productSellRequest.Quantity,
		Price:    p.Price,
	})

	fmt.Println("user basket: ", user.Basket)

	s.Profit += productSellRequest.Quantity * (p.Price - p.OriginalPrice)

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
