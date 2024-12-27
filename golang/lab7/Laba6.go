package LABA7

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
	dress := &Dress{"платье вечернее", 8000.00, "Zara", "красный"}
	wardrobe := &Wardrobe{"шкаф для одежды", 15000.00, "IKEA", "белый"}
	products := []Product{dress, wardrobe}
	fmt.Println("общая стоимость", Calculate(products))
	dress.ApplyDiscount(10)
	wardrobe.ApplyDiscount(5)
	fmt.Println("общая стоимость товаров после применения скидок", Calculate(products))
	fmt.Println("цвет платья:", dress.Color)
	fmt.Println("цвет шкафа:", wardrobe.Color)
	dress.ChangeColor("черный")
	wardrobe.ChangeColor("коричневый")
	fmt.Println("цвет платья изменен на:", dress.Color)
	fmt.Println("цвет шкафа изменен на:", wardrobe.Color)
}
