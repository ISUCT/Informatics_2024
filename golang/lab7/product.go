package lab7

type Product interface {
	GetPrice() float64
	ApplyingTheDiscount(float64)
	ChangePrice(float64)
	ChangeCharacteristic(string, string)
}

func CalculatePrice(list []Product) float64 {
	var sum float64 = 0
	for _, price := range list {
		sum += price.GetPrice()
	}
	return sum
}
