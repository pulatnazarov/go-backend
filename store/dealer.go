package main

type Dealer struct{}

func (d Dealer) ProvideProduct(productName string, budget, quantity int) (Product, bool) {
	originalPrice := generateProductPrice(1, 20) * 1000
	temp := 0

	if budget < originalPrice*quantity {
		temp = budget / originalPrice

		if temp < 1 {
			return Product{}, false
		}
	} else {
		temp = quantity
	}

	return Product{
		Name:          productName,
		Quantity:      temp,
		Price:         originalPrice * 2,
		OriginalPrice: originalPrice,
	}, true
}
