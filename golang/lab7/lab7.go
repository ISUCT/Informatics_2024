package lab7

import (
	"fmt"
)

type Product interface {
	GetName() string
	GetPrice() float64
	ApplyDiscount(discount float64)
}

func Calculate(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func RunLab7() {
	phone := &Phone{"Самсунг Колдснеп эс 18", 2000.00, "Самсунг", 30}
	tablet := &Tablet{"Эпл Торнадо Фордж 2", 4000.00, "Эпл", 40}
	products := []Product{phone, tablet}
	fmt.Println("Общая стоимость", Calculate(products))
	phone.ApplyDiscount(10)
	tablet.ApplyDiscount(15)
	fmt.Println("Стоимость товаров после скидок", Calculate(products))

	fmt.Println("Объём памяти телефона:", phone.Storage, "ГБ")
	fmt.Println("Объём памяти планшета:", tablet.Storage, "ГБ")
	phone.ChangeStorage(128)
	tablet.ChangeStorage(256)
	fmt.Println("Изменённый объём памяти телефона:", phone.Storage, "ГБ")
	fmt.Println("Изменённый объём памяти планшета:", tablet.Storage, "ГБ")
}
