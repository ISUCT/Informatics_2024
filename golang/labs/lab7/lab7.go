package lab7

import (
	"fmt"
)

func CalculateTotalPrice(products []Products) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.GetPrice()
	}

	return sum
}

func RunLab7() {
	product1 := NewTShirt("Nike", "White", 70, 1000)
	product2 := NewChips("Lays", "Crab", 100, 150)

	listOfProduct := []Products{product1, product2}

	fmt.Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%s-------%.2f ₽\n", product.GetName(), product.GetPrice())
	}
	fmt.Printf("Цена корзины до скидки: %.2f ₽\n\n", CalculateTotalPrice(listOfProduct))

	product1.ApplyDiscount(10)
	product2.ApplyDiscount(27)

	fmt.Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%s-------%.2f ₽\n", product.GetName(), product.GetPrice())
	}
	fmt.Printf("Цена корзины после скидки: %.2f ₽\n", CalculateTotalPrice(listOfProduct))

	fmt.Println("Информация про футболку")
	fmt.Println("Цвет:", product1.GetColor())
	fmt.Println("Размер:", product1.GetSize())
	fmt.Println("Бренд:", product1.GetName())

	fmt.Println("Информация про чипсы")
	fmt.Println("Вкус:", product2.GetTaste())
	fmt.Println("Масса:", product2.GetWeight())
	fmt.Println("Название:", product2.GetName())
}
