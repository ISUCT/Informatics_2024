package lab7

import "fmt"

type Book struct {
	name        string
	author      string
	price       float64
	description string
}

func NewBook(name string, author string, price float64, descrption string) *Book {
	book := &Book{name: name, author: author, price: price, description: descrption}
	return book
}

func (b *Book) getName() string {
	return b.name
}

func (b *Book) setPrice(newPrice float64) {
	b.price = newPrice
}

func (b *Book) getPrice() float64 {
	return b.price
}

func (b *Book) changeData(name string, price float64, descrption string) {
	b.name = name
	b.price = price
	b.description = descrption
}

func (b *Book) getData() {
	fmt.Printf("name: %s\nauthor: %s\nprice: %.2f\ndescrption: %s\n\n", b.name, b.author, b.price, b.description)
}

func (b *Book) applyDiscount(perDiscount float64) {
	b.price = b.price * (1 - (perDiscount / 100))
}
