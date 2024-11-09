package sevenlab

import (
	"fmt"
)

type Product interface {
	GetPrice() float64        // Получить цену продукта
	ApplyDiscount(float64)    // Применить скидку на продукт
	UpdatePrice(float64)      // Обновить цену продукта
	UpdateDescription(string) // Обновить описание продукта
}

type Book struct {
	price       float64
	description string
}

func (b *Book) GetPrice() float64 {
	return b.price
}

func (b *Book) ApplyDiscount(discount float64) {
	b.price = b.price * (1 - discount/100)
}

func (b *Book) UpdatePrice(newPrice float64) {
	b.price = newPrice
}

func (b *Book) UpdateDescription(newDescription string) {
	b.description = newDescription
}

type Electronics struct {
	price       float64
	description string
}

func (e *Electronics) GetPrice() float64 {
	return e.price
}

func (e *Electronics) ApplyDiscount(discount float64) {
	e.price = e.price * (1 - discount/100)
}

func (e *Electronics) UpdatePrice(newPrice float64) {
	e.price = newPrice
}

func (e *Electronics) UpdateDescription(newDescription string) {
	e.description = newDescription
}

func CalculateTotalPrice(products []Product) float64 {
	var total float64
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Start7lab() {
	book := &Book{price: 200, description: "Programming in Go"}
	electronics := &Electronics{price: 5000, description: "Smartphone"}

	products := []Product{book, electronics}

	totalBeforeDiscount := CalculateTotalPrice(products)
	fmt.Printf("Total price before discounts: %.2f\n", totalBeforeDiscount)

	book.ApplyDiscount(10)
	electronics.ApplyDiscount(15)

	totalAfterDiscount := CalculateTotalPrice(products)
	fmt.Printf("Total price after discounts: %.2f\n", totalAfterDiscount)

	book.UpdatePrice(180)
	book.UpdateDescription("Advanced Go Programming")
	electronics.UpdatePrice(4500)
	electronics.UpdateDescription("Smartphone Pro")

	finalTotal := CalculateTotalPrice(products)
	fmt.Printf("Final total price after updates: %.2f\n", finalTotal)
}
