package lab7

import "fmt"

type Product interface {
	GetPrice() float64
	EditPrice(float64)
	ApplyDiscount(float64)
	EditDescription(string)
	GetInfo()
}

func CalculatePrice(purchase []Product) float64 {
	var summ float64 = 0
	for _, price := range purchase {
		summ += price.GetPrice()
	}
	return summ
}

func RunLab7() {
	var car Product = &Car{"BMW", "чёрная", 300, 10000}
	var headphones Product = &Headphones{"Apple", "белые", 100, 25000}
	var keyboard Product = &Keyboard{"Red Square", "проводная", 86, 4599}

	car.GetInfo()
	headphones.GetInfo()
	keyboard.GetInfo()
	fmt.Println(car.GetPrice(), headphones.GetPrice(), keyboard.GetPrice())
	car.EditPrice(60)

	var purchase []Product = []Product{car, headphones, keyboard}
	summ := CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", summ)

	headphones.EditDescription("электронном")
	headphones.ApplyDiscount(40)
	keyboard.ApplyDiscount(60)
	keyboard.EditDescription("L")

	car.GetInfo()
	headphones.GetInfo()
	keyboard.GetInfo()

	summ = CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", summ)
}
