package LAB_7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

func Calculate(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Lab7() {
	laptop := &Laptop{"Делл XPS 13", 40000.00, "Dell", "Серый"}
	car := &Car{"Тойота Королла", 260000.00, "Toyota", "Красный"}

	products := []Product{laptop, car}
	fmt.Println("Общая стоимость", Calculate(products))

	laptop.ApplyDiscount(17)
	car.ApplyDiscount(7)
	fmt.Println("Общая стоимость товаров после применения скидок", Calculate(products))
}
