package lab7

import "fmt"

type Product interface {
	GetInfo()
	GetPrice() float64
	ApplyDiscount(float64)
	EditPrice(float64)
	EditDescription(string)
}

func CalculatePrice(products []Product) float64 {
	sum := 0.0
	for _, price := range products {
		sum += price.GetPrice()
	}
	return sum
}

func RunLab7Tasks() {
	var electronics Product = &Electronics{"Телефон", "Samsung", 24000}
	var vegetables Product = &Vegetables{"Картофель", "свежий", 60}
	var chancellery Product = &Chancellery{"Ручка", "гелевая", 25}
	electronics.GetInfo()
	vegetables.GetInfo()
	chancellery.GetInfo()
	var buy []Product = []Product{electronics, vegetables, chancellery}
	sum := CalculatePrice(buy)
	fmt.Println("Общая стоимость:", sum)

	vegetables.EditPrice(80)
	electronics.ApplyDiscount(30)
	chancellery.EditDescription("шариковая")
	vegetables.GetInfo()
	electronics.GetInfo()
	chancellery.GetInfo()

	sum = CalculatePrice(buy)
	fmt.Println("Общая стоимость:", sum)
}
