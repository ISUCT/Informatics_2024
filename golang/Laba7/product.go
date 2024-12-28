package lab7

import (
	"fmt"
)

type Product interface {
	GetName() string
	SetName(NewName string)
	GetPrice() float64
	SetPrice(NewPrice float64)
	ApplyDiscount(discount float64)
}

func CalculationSumProduct(listProducts []Product) string {
	var sum float64 = 0
	for _, product := range listProducts {
		sum += product.GetPrice()
	}
	s := fmt.Sprintf("%.2f", sum)
	return s
}
