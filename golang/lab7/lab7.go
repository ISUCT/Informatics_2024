package lab7

import "fmt"

type Product interface {
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
	GetName() string
	SetName(name string)
}

func CalculateTotalPrice(products []Product) float64 {
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += product.GetPrice()
	}
	return totalPrice
}

func Lab7() {
	iphone := &Phone{price: 100, weight: 200, name: "Айфон"}
	tshirt := &Shirt{price: 200, size: "S", name: "Футболка"}
	apple := &Fruits{price: 50, calories: 52, name: "Яблоко"}

	products := []Product{iphone, tshirt, apple}

	totalPrice := CalculateTotalPrice(products)
	fmt.Println("Общая цена без скидки:", totalPrice)

	products[0].ApplyDiscount(10)
	products[1].ApplyDiscount(20)
	products[2].ApplyDiscount(25)

	totalPrice = CalculateTotalPrice(products)
	fmt.Println("Общая цена с скидкой:", totalPrice)
}
