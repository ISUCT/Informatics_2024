package lab7

import (
	"fmt"
)

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

func CalculateTotalCost(products []Product) float64 {
	var total float64
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func RunLab7() {
	magazine := &Magazine{
		Name:      "Всё для дома",
		Publisher: "ЧитайГород",
		Price:     200,
	}

	car := &Car{
		Name:  "Ваз",
		Model: "VAZ2109",
		Price: 2500000,
	}

	products := []Product{magazine, car}

	totalBeforeDiscount := CalculateTotalCost(products)
	fmt.Printf("Общая стоимость до применения скидок: %.2f\n", totalBeforeDiscount)

	products[0].ApplyDiscount(10)
	products[1].ApplyDiscount(5)

	totalAfterDiscount := CalculateTotalCost(products)
	fmt.Printf("Общая стоимость после применения скидок: %.2f\n", totalAfterDiscount)

	car.SetModel("VAZ3109")
	magazine.SetName("Всё для дома")

	fmt.Println(car)
	fmt.Println(magazine)
}
