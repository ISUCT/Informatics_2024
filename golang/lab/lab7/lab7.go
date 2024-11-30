package lab7

import (
	"fmt"
)

type Product interface {
	GetInfo()
	GetPrice() float32
	MakeDiscount(float32)
	ChangePrice(float32)
	ChangeDescription(string)
}

func CalculatePrice(list []Product) float32 {
	var sum float32 = 0
	for _, price := range list {
		sum += price.GetPrice()
	}
	return sum
}

func RunLab7Tasks() {
	var fruit Product = &Fruit{"яблоко", "свежее", 40}
	var book Product = &Book{"Война и мир", "бумажном", 500}
	var clothes Product = &Clothes{"худи", "M", 3500}
	fruit.GetInfo()
	book.GetInfo()
	clothes.GetInfo()
	var purchase []Product = []Product{fruit, book, clothes}
	sum := CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", sum)

	fruit.ChangePrice(60)
	book.ChangeDescription("электронном")
	book.MakeDiscount(40)
	clothes.MakeDiscount(60)
	clothes.ChangeDescription("L")
	fruit.GetInfo()
	book.GetInfo()
	clothes.GetInfo()

	sum = CalculatePrice(purchase)
	fmt.Println("Общая стоимость товара:", sum)
}
