package labs

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
	book1 := Book{title: "The Go Programming Language", author: "Alan Donovan", price: 40.0}
	book2 := Book{title: "Programming in Go", author: "Mark Summerfield", price: 35.0}
	electronics1 := Electronics{name: "iPhone 13", brand: "Apple", price: 1000.0}

	products := []Product{&book1, &book2, &electronics1}

	fmt.Println("Initial Total Price:", CalculateTotalPrice(products))

	book1.ApplyDiscount(10)
	electronics1.SetPrice(900.0)

	fmt.Println("Total Price after discounts:", CalculateTotalPrice(products))

	for _, product := range products {
		fmt.Println(product.GetDetails())
	}
}
