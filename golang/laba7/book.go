package laba7

import "fmt"

type Book struct {
	name   string
	price  float64
	author string
	genre  string
}

func (b *Book) applyDiscount(discount float64) {
	b.price = b.price * (1 - discount/100)
}

func (b *Book) getPrice() float64 {
	return b.price
}

func (b *Book) getProductInfo() string {
	return fmt.Sprintf("Name: %s, Author: %s, Genre: %s, Price: %.2f", b.name, b.author, b.genre, b.price)
}
