package lab7

import (
	"fmt"
)

func CalculateTotal(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Executelab7() {
	Laptop := &Laptop{Name: "Laptop", Price: 1000}
	Phone := &Phone{Name: "Phone", Price: 500}
	Tablet := &Tablet{Name: "Tablet", Price: 300}

	products := []Product{Laptop, Phone, Tablet}

	fmt.Printf("Total before discount: %.2f\n", CalculateTotal(products))

	for _, product := range products {
		product.ApplyDiscount(10)
	}

	fmt.Printf("Total after discount: %.2f\n", CalculateTotal(products))
}
