package Lab_7

import (
	"fmt"
)

type Product interface {
	GetPrice() float32
	Discount(float32)
	ChangePrice(float32)
	ChangeCharacteristic(string, string)
}

func CalculateTotalPrice(products []Product) float32 {
	var total float32
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func RunLab7() {
	orange := &Orange{price: 40000, color: "оранжевый", material: "пластик"}
	apple := &Apple{price: 17000, color: "красный", material: "металл"}
	bike := &Bike{price: 9000, color: "зелёный", material: "алюминий"}
	roller := &Roller{price: 12000, color: "синий", material: "пластик"}

	products := []Product{orange, apple, bike, roller}

	fmt.Println("Общая стоимость товаров:", CalculateTotalPrice(products), "рублей")

	orange.Discount(25)
	apple.Discount(50)
	bike.Discount(10)
	roller.Discount(20)

	orange.ChangeCharacteristic("жёлтый", "дерево")
	apple.ChangeCharacteristic("зелёный", "пластик")
	
	fmt.Println("Общая стоимость товаров после скидок и изменений:", CalculateTotalPrice(products), "рублей")
}
