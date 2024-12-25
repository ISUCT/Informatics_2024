package lab7

import "fmt"

type TShirt struct {
	name  string
	color string
	size  float64
	price float64
}

func NewTShirt(name string, color string, size float64, price float64) *TShirt {
	tshirt := &TShirt{name: name, color: color, price: price, size: size}
	return tshirt
}

func (t *TShirt) GetName() string {
	return t.name
}

func (t *TShirt) SetPrice(newPrice float64) {
	t.price = newPrice
}

func (t *TShirt) GetPrice() float64 {
	return t.price
}

func (t *TShirt) GetSize() float64 {
	return t.size
}

func (t *TShirt) GetColor() string {
	return t.color
}

func (t *TShirt) GetInfo() {
	fmt.Printf("name: %s\ncolor: %s\nsize: %.2f\nprice: %.2f\n", t.name, t.color, t.size, t.price)
}

func (t *TShirt) ApplyDiscount(perDiscount float64) {
	t.price = t.price * (1 - (perDiscount / 100))
}
