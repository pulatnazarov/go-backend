package dealer

import (
	"go-backend/store/product"
)

type Dealer struct{}

func (d Dealer) ProvideProduct(productName string, budget, quantity int) (product.Product, bool) {
	originalPrice := product.GenerateProductPrice(1, 20) * 1000
	temp := 0

	if budget < originalPrice*quantity {
		temp = budget / originalPrice

		if temp < 1 {
			return product.Product{}, false
		}
	} else {
		temp = quantity
	}

	return product.Product{
		Name:          productName,
		Quantity:      temp,
		Price:         originalPrice * 2,
		OriginalPrice: originalPrice,
	}, true
}
