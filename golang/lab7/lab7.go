package lab7

import (
	"fmt"
)

type Product interface {
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
	GetDetails() string
}

type Car struct {
	Model string
	Brand string
	Price float64
}

func (b *Car) GetPrice() float64 {
	return b.Price
}

func (b *Car) SetPrice(price float64) {
	b.Price = price
}

func (b *Car) ApplyDiscount(discount float64) {
	b.Price = b.Price * (1 - discount/100)
}

func (b *Car) GetDetails() string {
	return fmt.Sprintf("Car: %s by %s, Price: %.2f", b.Model, b.Brand, b.Price)
}

type Electronic struct {
	Name  string
	Brand string
	Price float64
}

func (e *Electronic) GetPrice() float64 {
	return e.Price
}

func (e *Electronic) SetPrice(price float64) {
	e.Price = price
}

func (e *Electronic) ApplyDiscount(discount float64) {
	e.Price = e.Price * (1 - discount/100)
}

func (e *Electronic) GetDetails() string {
	return fmt.Sprintf("Electronic: %s by %s, Price: %.2f", e.Name, e.Brand, e.Price)
}

type Book struct {
	Title  string
	Author string
	Price  float64
}

func (b *Book) GetPrice() float64 {
	return b.Price
}

func (b *Book) SetPrice(price float64) {
	b.Price = price
}

func (b *Book) ApplyDiscount(discount float64) {
	b.Price = b.Price * (1 - discount/100)
}

func (b *Book) GetDetails() string {
	return fmt.Sprintf("Book: %s by %s, Price: %.2f", b.Title, b.Author, b.Price)
}

func CalculateTotalCost(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Run7() {
	car1 := &Car{Model: "X Class 350d", Brand: "Mercedes", Price: 199999.00}
	car2 := &Car{Model: "S Class 450 Long", Brand: "Mercedes", Price: 300000.00}
	electronic1 := &Electronic{Name: "Iphone 16", Brand: "Apple", Price: 1200.00}
	electronic2 := &Electronic{Name: "ASUS Tuf Gaming A15", Brand: "Asus", Price: 999.99}
	book1 := &Book{Title: "Learn C Programming", Author: "Jeff Szuhay", Price: 12.00}
	book2 := &Book{Title: "Go Systems Programming", Author: "Mihalis Tsoukalos", Price: 10.99}

	products := []Product{car1, car2, electronic1, electronic2, book1, book2}

	fmt.Println("Товары до получения скидок:")
	for _, product := range products {
		fmt.Println(product.GetDetails())
	}
	fmt.Printf("Общая стоимость без учета скидок: %.2f\n", CalculateTotalCost(products))

	car1.ApplyDiscount(10)
	car2.ApplyDiscount(5)
	electronic1.ApplyDiscount(20)
	book2.ApplyDiscount(5)

	fmt.Println("\nТовары после скидок и изменения цен:")
	for _, product := range products {
		fmt.Println(product.GetDetails())
	}
	fmt.Printf("Общая стоимость после учета скидок: %.2f\n", CalculateTotalCost(products))
}
