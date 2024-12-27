package lab7

import (
	"fmt"
)

type Product interface {
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
	GetName() string
}

func CalculateTotalPrice(products []Product) float64 {
	var total float64
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Runlab7() {
	brush := &Brush{name: "Иммитация колонка", price: 560.00, material: "Колонок"}
	palette := &Palette{name: "Деревянная палитра", price: 160.00, colorCount: 12}
	canvas := &Canvas{name: "Грунтованный холст", price: 300.00, size: "100x80cm"}

	products := []Product{brush, palette, canvas}

	fmt.Println("Сумма товаров без учёта скидки:", CalculateTotalPrice(products))
	products[0].ApplyDiscount(10)
	products[1].ApplyDiscount(15)
	products[2].ApplyDiscount(20)
	fmt.Println("Сумма товаров с учётом скидки:", CalculateTotalPrice(products))
}
