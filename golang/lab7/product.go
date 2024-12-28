package lab7

type Product interface {
	setdiscount(discount float64)
	setPrice(newPrice float64)
	getPrice() float64
	getName() string
	GetInfo() string
}

func CalculateProductsSum(products []Product) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.getPrice()
	}
	return sum
}
