package main

type Repository struct {
	Products ProductList
}

func (r Repository) NewRepository(products ProductList) Repository {
	return Repository{
		Products: products,
	}
}

func (r Repository) Search(productName string) (Product, bool) {
	for _, product := range r.Products {
		if product.Name == productName {
			return product, true
		}
	}

	return Product{}, false
}

func (r Repository) TakeProduct(productName string, quantity int) {
	for i, product := range r.Products {
		if product.Name == productName {
			r.Products[i].Quantity -= quantity
			return
		}
	}
}
