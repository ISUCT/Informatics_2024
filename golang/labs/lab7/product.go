package lab7

type Product interface {
	make_sale(discount float64)
	get_price() float64
	get_productInfo() string
}

func get_priceAllProducts(products []Product) float64 {
	var totalCost float64
	for _, product := range products {
		totalCost += product.get_price()
	}
	return totalCost
}
