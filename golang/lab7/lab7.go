package lab7

import "fmt"

type Product interface {
	GetInfo()
	GetPrice() float64
	MakeDiscount(float64)
	ChangePrice(float64)
	ChangeDescription(string)
}

func CalculatePrice(list []Product) float64 {
	var sum float64 = 0
	for _, item := range list {
		sum += item.GetPrice()
	}
	return sum
}

func RunLab7Tasks() {
	shoes := &Shoes{"кроссовки", "38", 50}
	book := &Book{"Преступление и наказание", "бумажном", 700}
	car := &Car{"машина", "БМВ", 10000}
	shoes.GetInfo()
	book.GetInfo()
	car.GetInfo()
	var purchase = []Product{shoes, book, car}
	sum := CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", sum)
	shoes.ChangePrice(60)
	book.ChangeDescription("электронном")
	book.MakeDiscount(20) // Скидка 20%
	car.MakeDiscount(10) // Скидка 10%
    car.ChangeDescription("Мерседес")
	shoes.GetInfo()
	book.GetInfo()
	car.GetInfo()
	sum = CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", sum)
}

