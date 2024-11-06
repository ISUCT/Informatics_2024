package lab7

import (
	"fmt"
)

// Интерфейс Product
type Product interface {
	GetPrice() float64
	GetName() string
	ApplyDiscount(discount float64)
	SetPrice(price float64)
	SetName(name string)
	GetDescription() string
	SetDescription(description string)
}

// Структура для электроники
type Electronic struct {
	name        string
	price       float64
	description string
}

// Реализация интерфейса Product для Electronic
func (e *Electronic) GetPrice() float64 {
	return e.price
}

func (e *Electronic) GetName() string {
	return e.name
}

func (e *Electronic) ApplyDiscount(discount float64) {
	e.price -= e.price * discount
}

func (e *Electronic) SetPrice(price float64) {
	e.price = price
}

func (e *Electronic) SetName(name string) {
	e.name = name
}

func (e *Electronic) GetDescription() string {
	return e.description
}

func (e *Electronic) SetDescription(description string) {
	e.description = description
}

// Структура для одежды
type Clothing struct {
	name        string
	price       float64
	description string
	size        string
}

// Реализация интерфейса Product для Clothing
func (c *Clothing) GetPrice() float64 {
	return c.price
}

func (c *Clothing) GetName() string {
	return c.name
}

func (c *Clothing) ApplyDiscount(discount float64) {
	c.price -= c.price * discount
}

func (c *Clothing) SetPrice(price float64) {
	c.price = price
}

func (c *Clothing) SetName(name string) {
	c.name = name
}

func (c *Clothing) GetDescription() string {
	return c.description
}

func (c *Clothing) SetDescription(description string) {
	c.description = description
}

// Функция для расчета стоимости списка товаров
func CalculateTotalPrice(products []Product) float64 {
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += product.GetPrice()
	}
	return totalPrice
}

func RunLab7() {
	// Пример использования
	phone := &Electronic{name: "iPhone 14", price: 1000, description: "New iPhone"}
	tshirt := &Clothing{name: "T-Shirt", price: 20, description: "Cotton T-Shirt", size: "M"}

	products := []Product{phone, tshirt}

	fmt.Println("Total price before discount:", CalculateTotalPrice(products))

	phone.ApplyDiscount(0.1)   // 10% скидка на телефон
	tshirt.ApplyDiscount(0.05) // 5% скидка на футболку

	fmt.Println("Total price after discount:", CalculateTotalPrice(products))

	//Пример изменения характеристик
	tshirt.SetName("New T-Shirt")
	tshirt.SetPrice(25)
	fmt.Println("Price of T-Shirt after changes:", tshirt.GetPrice())
	fmt.Println("Name of T-Shirt after changes:", tshirt.GetName())

}
