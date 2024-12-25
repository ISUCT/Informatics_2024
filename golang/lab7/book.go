package lab7

import "fmt"

type Book struct {
	name string
	author string
	price float64
	description string
}

func NewBook(name string, author string, price float64, description string) *Book {
	book := &Book{name: name, author: author, price: price, description: description}
	return book
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) SetPrice(newPrice float64) {
	b.price = newPrice
}

func (b *Book) GetPrice() float64 {
	return b.price
}

func (b *Book) ChangeData(name string, price float64, description string) {
	b.name = name
	b.price = price
	b.description = description
}

func (b *Book) GetData() {
	fmt.Printf("name: %s\nauthor: %s\nprice: %.2f\ndescription: %s\n\n", b.name, b.author, b.price, b.description)
}

func (b *Book) ApplyDiscount(perDiscount float64) {
	b.price = b.price * (1 - (perDiscount / 100))
}