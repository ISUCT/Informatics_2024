package lab7

func CalculateTotalPrice(products []Product) float64 {
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += product.GetPrice()
	}
	return totalPrice
}
