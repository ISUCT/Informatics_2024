package main

import (
	"fmt"
)

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
