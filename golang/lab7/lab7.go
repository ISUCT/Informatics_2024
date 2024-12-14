package lab7

import (
	"fmt"
)

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

func CalculateTotalCost(products []Product) float64 {
	var total float64
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Runlab7() {
	book := &Book{
		Name:   "Убить пересмешника",
		Author: "Харпер Ли",
		Price:  500,
	}

	electronic := &Electronics{
		Name:  "Смартфон",
		Model: "iPhone 14 Pro Max",
		Price: 120000,
	}

	products := []Product{book, electronic}

	totalBeforeDiscount := CalculateTotalCost(products)
	fmt.Printf("Общая стоимость до применения скидок: %.2f\n", totalBeforeDiscount)

	book.ApplyDiscount(10)
	electronic.ApplyDiscount(5)

	totalAfterDiscount := CalculateTotalCost(products)
	fmt.Printf("Общая стоимость после применения скидок: %.2f\n", totalAfterDiscount)
}
