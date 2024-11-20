package lab

import (
	"fmt"
)

type Product interface {
	GetInfo()
	GetPrice() float32
	Sale(float32)
	ChangePrice(float32)
	ChangeChar(string)
}

func CalculatePrice(list []Product) float32 {
	var sum float32 = 0
	for _, price := range list {
		sum += price.GetPrice()
	}
	return sum
}

func RunLab7Tasks() {
	var interface_fruit Product = &Fruit{"яблоко", "свежее", 40}
	var interface_book Product = &Book{"Война и мир", "бумажном", 500}
	var interface_clothes Product = &Clothes{"худи", "M", 3500}
	interface_fruit.GetInfo()
	interface_book.GetInfo()
	interface_clothes.GetInfo()
	var purchase []Product = []Product{interface_fruit, interface_book, interface_clothes}
	sum := CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", sum)

	interface_fruit.ChangePrice(60)
	interface_book.ChangeChar("электронном")
	interface_book.Sale(40)
	interface_clothes.Sale(60)
	interface_clothes.ChangeChar("L")
	interface_fruit.GetInfo()
	interface_book.GetInfo()
	interface_clothes.GetInfo()

	sum = CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", sum)
}
