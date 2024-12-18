package laba7

import "fmt"

type Book struct {
	Name   string
	Price  float64
	Author string
	Genre  string
}

func (b *Book) ApplyDiscount(discount float64) {
	b.Price = b.Price * (1 - discount/100)
}

func (b *Book) GetPrice() float64 {
	return b.Price
}

func (b *Book) GetProductInfo() string {
	return fmt.Sprintf("Name: %s, Author: %s, Genre: %s, Price: %.2f", b.Name, b.Author, b.Genre, b.Price)
}
