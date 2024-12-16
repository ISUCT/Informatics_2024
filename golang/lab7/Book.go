package lab7

import (
	"fmt"
)

type Book struct {
	Name   string
	Author string
	Price  float64
}

func NewBook(name string, author string, price float64) *Book {
	b := new(Book)
	b.Name = name
	b.Author = author
	b.Price = price
	return b
}

func (b *Book) SetPrice(price float64)  { b.Price = price }
func (b Book) GetPrice() float64        { return b.Price }
func (b *Book) SetName(name string)     { b.Name = name }
func (b Book) GetName() string          { return b.Name }
func (b *Book) SetAuthor(author string) { b.Author = author }

func (b *Book) ApplyDiscount(discount float64) {
	if discount > 0 && discount <= 100 {
		b.SetPrice(b.GetPrice() * (1 - discount/100))
	} else {
		fmt.Println("Неверная скидка")
	}
}
