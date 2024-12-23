package lab7

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

func CalculateTotalPrice(products []Product) float64 {
	var total float64
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}
