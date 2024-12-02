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

type Item struct {
	Name  string
	Price float64
}

func (i *Item) GetName() string {
	return i.Name
}

func (i *Item) GetPrice() float64 {
	return i.Price
}

func (i *Item) SetPrice(price float64) {
	i.Price = price
}

func (i *Item) ApplyDiscount(discount float64) {
	i.Price -= discount
	if i.Price < 0 {
		i.Price = 0
	}
}

func CalculateTotal(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func RunLab7() {    
	item1 := &Item{Name: "Утюг", Price: 1000.0}
	item2 := &Item{Name: "Микроволновка", Price: 1500.0}
	item3 := &Item{Name: "Холодильник", Price: 3000.0}

	item1.ApplyDiscount(100)
	item2.ApplyDiscount(50)

	products := []Product{item1, item2, item3}

	fmt.Printf("Общая стоимость без учёта скидок: %.2f\n", CalculateTotal([]Product{
		&Item{Name: "Утюг", Price: 1000.0},
		&Item{Name: "Микроволновка", Price: 1500.0},
		&Item{Name: "Холодильник", Price: 3000.0},
	}))

	totalAfterDiscounts := CalculateTotal(products)
	fmt.Printf("Общая стоимость с учётом скидок: %.2f\n", totalAfterDiscounts)
}
