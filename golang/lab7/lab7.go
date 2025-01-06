package lab7

import "fmt"

type Product interface {
	GetDescription()
	GetPrice() float64
	ApplyDiscount(discount float64)
	SetNewPrice(newPrice float64)
	SetNewDescription(newDescription string)
}

func CalculatePrice(list []Product) float64 {
	var sum float64 = 0
	for _, price := range list {
		sum += price.GetPrice()
	}
	return sum
}

func RunLab7() {
	var CatFood Product = &CatFood{"Orijen Cat&Kitten Grin Free", "Orijen", 9900}
	var Microwave_oven Product = &Microwave_oven{"Hotpoint MWHA 201", 10490}
	var Dress Product = &Dress{"гипюровое платье", "полиэстера", "DStrend", 5983}

	CatFood.GetDescription()
	Microwave_oven.GetDescription()
	Dress.GetDescription()

	var purchase []Product = []Product{CatFood, Microwave_oven, Dress}
	sum := CalculatePrice(purchase)
	fmt.Println("Общая стоимость списка товаров:", sum)

	CatFood.SetNewPrice(9400)
	Microwave_oven.SetNewDescription("Samsung")
	Microwave_oven.ApplyDiscount(0.1)
	Dress.ApplyDiscount(0.2)
	Dress.SetNewDescription("хлопка")

	CatFood.GetDescription()
	Microwave_oven.GetDescription()
	Dress.GetDescription()

	sum = CalculatePrice(purchase)
	fmt.Println("Общая стоимость списка товаров с учетом всех скидок:", sum)
}
