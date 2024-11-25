package lab7

import (
	"fmt"
	"math/rand"
	"time"
)

type Product interface {
	GetPrice() float64
	SetPrice(float64)
	ApplyDiscount(float64)
	String() string
}

type Book struct {
	Title     string
	Author    string
	Publisher string
	Price     float64
}

func (b *Book) GetPrice() float64              { return b.Price }
func (b *Book) SetPrice(price float64)         { b.Price = price }
func (b *Book) ApplyDiscount(discount float64) { b.Price *= (1 - discount/100) }
func (b *Book) String() string {
	return fmt.Sprintf("Книга '%s' автора '%s', издательство '%s', цена: %.2f", b.Title, b.Author, b.Publisher, b.Price)
}

type Electronics struct {
	Name  string
	Brand string
	Model string
	Price float64
}

func (e *Electronics) GetPrice() float64              { return e.Price }
func (e *Electronics) SetPrice(price float64)         { e.Price = price }
func (e *Electronics) ApplyDiscount(discount float64) { e.Price *= (1 - discount/100) }
func (e *Electronics) String() string {
	return fmt.Sprintf("Электроника '%s' бренда '%s', модель '%s', цена: %.2f", e.Name, e.Brand, e.Model, e.Price)
}

func CalculateTotalCost(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Complitelab7() {
	rand.Seed(time.Now().UnixNano())

	book1 := &Book{"Властелин колец", "Дж. Р. Р. Толкин", "HarperCollins", 25.99}
	book2 := &Book{"1984", "Джордж Оруэлл", "Penguin Books", 12.95}
	electronic1 := &Electronics{"iPhone XR", "Apple", "XR", 749.00}
	electronic2 := &Electronics{"MacBook Pro", "Apple", "16-inch, M1 Max", 2499.00}

	products := []Product{book1, book2, electronic1, electronic2}

	book1.ApplyDiscount(25)
	electronic1.ApplyDiscount(22)

	totalCost := CalculateTotalCost(products)

	fmt.Println("Список товаров:")
	for _, p := range products {
		fmt.Println(p.String())
	}

	fmt.Printf("\nОбщая стоимость: $%.2f\n", totalCost)
}
