package lab7

import (
	"fmt"
)

type Product interface {
	GetInfo()
	GetPrice() float64
	SetDiscount(float64)
	EditPrice(float64)
	EditDescription(string)
}

func CalculatePrice(purchase []Product) float64 {
	var total float64 = 0
	for _, price := range purchase {
		total += price.GetPrice()
	}
	return total
}

func RunLab7() {
	var drink Product = &Drink{"кола", "1.95 л", 160}
	var toy Product = &Toy{"котик", "25 см", 1650}
	var lipstick Product = &Lipstick{"NARS", "бардовая", 4700}
	var purchase []Product = []Product{drink, toy, lipstick}

	drink.GetInfo()
	toy.GetInfo()
	lipstick.GetInfo()

	total := CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", total)

	drink.EditPrice(190)
	toy.EditDescription("1 м")
	toy.SetDiscount(10)
	lipstick.SetDiscount(5)
	lipstick.EditDescription("алая")

	drink.GetInfo()
	toy.GetInfo()
	lipstick.GetInfo()

	total = CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", total)
}
