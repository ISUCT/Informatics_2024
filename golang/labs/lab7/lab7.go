package lab7

import (
	"fmt"
)

type Product interface {
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
	GetName() string
}
type Brush struct {
	name     string
	price    float64
	material string
}

func (b *Brush) GetName() string {
	return b.name
}
func (b *Brush) GetPrice() float64 {
	return b.price
}
func (b *Brush) SetPrice(price float64) {
	b.price = price
}
func (b *Brush) ApplyDiscount(discount float64) {
	b.price -= b.price * discount / 100
}

type Palette struct {
	name       string
	price      float64
	colorCount int
}

func (p *Palette) GetName() string {
	return p.name
}

func (p *Palette) GetPrice() float64 {
	return p.price
}

func (p *Palette) SetPrice(price float64) {
	p.price = price
}

func (p *Palette) ApplyDiscount(discount float64) {
	p.price -= p.price * discount / 100
}

type Canvas struct {
	name  string
	price float64
	size  string
}

func (c *Canvas) GetName() string {
	return c.name
}

func (c *Canvas) GetPrice() float64 {
	return c.price
}

func (c *Canvas) SetPrice(price float64) {
	c.price = price
}

func (c *Canvas) ApplyDiscount(discount float64) {
	c.price -= c.price * discount / 100
}

func CalculateTotalPrice(products []Product) float64 {
	var total float64
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Runlab7() {
	brush := &Brush{name: "Иммитация колонка", price: 560.00, material: "Колонок"}
	palette := &Palette{name: "Деревянная палитра", price: 160.00, colorCount: 12}
	canvas := &Canvas{name: "Грунтованный холст", price: 300.00, size: "100x80cm"}

	products := []Product{brush, palette, canvas}

	fmt.Println("Сумма товаров без учёта скидки:", CalculateTotalPrice(products))
	brush.ApplyDiscount(10)
	palette.ApplyDiscount(15)
	canvas.ApplyDiscount(20)
	fmt.Println("Сумма товаров с учётом скидки:", CalculateTotalPrice(products))
}
