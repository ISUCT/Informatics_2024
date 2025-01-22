package main

import (
	"fmt"
)

type Product interface {
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(percent float64)
	GetDetails() string
}

type Book struct {
	title  string
	author string
	price  float64
}

func (b *Book) GetPrice() float64 {
	return b.price
}

func (b *Book) SetPrice(price float64) {
	b.price = price
}

func (b *Book) ApplyDiscount(percent float64) {
	b.price -= b.price * percent / 100
}

func (b *Book) GetDetails() string {
	return fmt.Sprintf("Book: %s, Author: %s, Price: %.2f", b.title, b.author, b.price)
}

type Electronics struct {
	name  string
	brand string
	price float64
}

func (e *Electronics) GetPrice() float64 {
	return e.price
}

func (e *Electronics) SetPrice(price float64) {
	e.price = price
}

func (e *Electronics) ApplyDiscount(percent float64) {
	e.price -= e.price * percent / 100
}
func (e *Electronics) GetDetails() string {
	return fmt.Sprintf("Electronics: %s, Brand: %s, Price: %.2f", e.name, e.brand, e.price)
}

func CalculateTotalPrice(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func main() {
	book := Book{title: "The Go Programming Language", author: "Alan Donovan", price: 50.00}
	electronics := Electronics{name: "Laptop", brand: "XYZ", price: 1200.00}

	fmt.Println(book.GetDetails())
	fmt.Println(electronics.GetDetails())

	products := []Product{&book, &electronics}
	totalPrice := CalculateTotalPrice(products)
	fmt.Printf("Общая стоимость: %.2f\n", totalPrice)
	book.ApplyDiscount(10)
	fmt.Println("Цена со скидкой:", book.GetDetails())

}
