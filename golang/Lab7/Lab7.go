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

func RunLab7A() {
	item1 := &Iron{Item: Item{Name: "Утюг", Price: 1000.0}}
	item2 := &Microwave{Item: Item{Name: "Микроволновка", Price: 1500.0}}
	item3 := &Refrigerator{Item: Item{Name: "Холодильник", Price: 3000.0}}

	item1.ApplyDiscount(100)
	item2.ApplyDiscount(50)

	products := []Product{item1, item2, item3}

	fmt.Printf("Общая стоимость без учёта скидок: %.2f\n", CalculateTotal([]Product{
		&Iron{Item: Item{Name: "Утюг", Price: 1000.0}},
		&Microwave{Item: Item{Name: "Микроволновка", Price: 1500.0}},
		&Refrigerator{Item: Item{Name: "Холодильник", Price: 3000.0}},
	}))

	totalAfterDiscounts := CalculateTotal(products)
	fmt.Printf("Общая стоимость с учётом скидок: %.2f\n", totalAfterDiscounts)
}

func RunLab7() {
	RunLab7()
}
