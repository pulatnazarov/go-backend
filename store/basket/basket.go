package basket

import "go-backend/store/product"

type Basket struct {
	ProductList product.ProductList
	Total       int
}

func NewBasket() Basket {
	return Basket{}
}
