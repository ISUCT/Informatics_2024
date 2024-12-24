package Lab7

import (
	"fmt"
)

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

func CalculateTotal(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func RunLab7() {
	printer := &Printer{Item: Item{Name: "Принтер", Price: 5000.0}}
	mouse := &Mouse{Item: Item{Name: "Компьютерная мышь", Price: 1500.0}}
	laptop := &Laptop{Item: Item{Name: "Ноутбук", Price: 30000.0}}

	printer.ApplyDiscount(500)
	mouse.ApplyDiscount(200)

	products := []Product{printer, mouse, laptop}

	fmt.Printf("Общая стоимость без учёта скидок: %.2f\n", CalculateTotal([]Product{
		&Printer{Item: Item{Name: "Принтер", Price: 5000.0}},
		&Mouse{Item: Item{Name: "Компьютерная мышь", Price: 500.0}},
		&Laptop{Item: Item{Name: "Ноутбук", Price: 30000.0}},
	}))

	totalAfterDiscounts := CalculateTotal(products)
	fmt.Printf("Общая стоимость с учётом скидок: %.2f\n", totalAfterDiscounts)
}
