package lab7

import (
	"fmt"
)

type Product interface {
	GetInfo()
	GetPrice() float64
	EditPrice(float64)
	ApplyDiscount(float64)
	EditDescription(string)
}

func CalculatePrice(list []Product) float64 {
	var sum float64 = 0
	for _, price := range list {
		sum += price.GetPrice()
	}
	return sum
}

func RunLab7() {
	var mouse Product = &Mouse{"Razer", "оптическим светодиодным", 7499}
	var keyboard Product = &Keyboard{"HyperX", "беспроводным", 8999}
	var headphones Product = &Headphones{"Logitech", "беспроводным", 12999}

	mouse.GetInfo()
	keyboard.GetInfo()
	headphones.GetInfo()

	var purchase []Product = []Product{mouse, keyboard, headphones}
	sum := CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", sum)

	mouse.EditDescription("оптическим")
	mouse.ApplyDiscount(15)
	keyboard.EditDescription("проводным")
	keyboard.ApplyDiscount(20)
	headphones.EditDescription("проводным")
	headphones.ApplyDiscount(10)

	mouse.GetInfo()
	keyboard.GetInfo()
	headphones.GetInfo()

	sum = CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", sum)
}
