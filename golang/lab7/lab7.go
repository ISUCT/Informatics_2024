package lab7

import (
	"fmt"
)

type Product interface {
	GetPrice() float32
	Sale(float32)
	ChangePrice(float32)
	changeCharacteristic(string, string)
}
func CalculatePrice(list []Product) float32 {
	var sum float32 = 0
	for _, price := range list {
		sum += price.GetPrice()
	}
	return sum
}

func RunLab7() {
	product1 := &Bicycles{40000, "черный", "аллюминиевый"}
	product2 := &Scooters{17000, "серебристый", "металлический"}
	product3 := &Skateboards{9000, "зеленый", "деревянный"}
	products := []Product{product1, product2, product3}
	fmt.Println("Общая стоимость товаров:", CalculatePrice(products), "рублей")
	product1.Sale(25)
	product2.Sale(50)
	product3.Sale(10)
	fmt.Println("Общая стоимость товаров после применения скидок:", CalculatePrice(products), "рублей")
}