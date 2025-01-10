package lab7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

func CalculatePrice(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func CompleteLab7() {
	vegetable := &vegetable{"Cucumber", 13000, "Хрустящие"}
	car := &Car{"Lada", 10000000, "Красный"}
	products := []Product{vegetable, car}
	fmt.Println("Общая стоимость", CalculatePrice(products))
	vegetable.ApplyDiscount(50)
	car.ApplyDiscount(2)
	fmt.Println("Общая стоимость после скидки", CalculatePrice(products))
	fmt.Println("Цвет машины:", car.Color)
	car.ChangeColor("Малиновая")
	fmt.Println("Да, у нас есть для вас ", car.Color,"Лада")
    fmt.Println("Сорт огурчиков:",vegetable.Sort)
    vegetable.GetSort("Малиновка")
    fmt.Println("Ваш сорт:",vegetable.Sort)
}