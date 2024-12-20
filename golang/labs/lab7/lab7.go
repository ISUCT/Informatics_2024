package lab7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	GetInfo()
	ApplyDiscount(discount float64)
}

func CalculateTotalPrice(products []Product) float64 {
	var sum float64 = 0.0
	for _, product := range products {
		sum += product.GetPrice()
	}
	return sum
}

func RunLab7() {
	product1 := NewIcecream("Магнат", 89.99, "Шоколад")
	product2 := NewMilk("Весёлая коровка", 79.99, "3.2")

	listOfProduct := []Product{product1, product2}

	fmt.Println("Товар-Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%v-%.2f $\n", product.GetName(), product.GetPrice())
	}
	fmt.Printf("Цена: %.2f $\n\n", CalculateTotalPrice(listOfProduct))

	product1.ApplyDiscount(5)
	product2.ApplyDiscount(10)

	fmt.Printf("Цена со скидкой: %.2f $\n", CalculateTotalPrice(listOfProduct))

	fmt.Println("Информация о мороженном")
	fmt.Println("Марка:", product1.GetName())
	fmt.Println("Цена:", product1.GetPrice())
	fmt.Println("Вкус:", product1.GetTaste())

	fmt.Println("Информация про молоко")
	fmt.Println("Марка:", product2.GetName())
	fmt.Println("Процент жирности:", product2.GetPercent())
}
