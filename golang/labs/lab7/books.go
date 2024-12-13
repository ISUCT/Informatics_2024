package lab7

import "fmt"

type Book struct {
	name          string
	author        string
	price         float64
	description   string
	numberOfPages int32
}

func NewBook(name string, author string, price float64, descrption string, numberOfPages int32) *Book {
	book := &Book{name: name, author: author, price: price, description: descrption, numberOfPages: numberOfPages}
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

func (b *Book) getAuthor() string {
	return b.author
}

func (b *Book) getNumberOfPages() int32 {
	return b.numberOfPages
}

func (b *Book) getInfo() {
	fmt.Printf("name: %s\nnumber of pages: %d\nauthor: %s\nprice: %.2f\ndescrption: %s\n\n", b.name, b.numberOfPages, b.author, b.price, b.description)
}

func (b *Book) applyDiscount(perDiscount float64) {
	b.price = b.price * (1 - (perDiscount / 100))
}
