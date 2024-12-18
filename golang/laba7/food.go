package laba7

import (
	"fmt"
)

type Food struct {
	name   string
	price  float64
	weight float64
}

func (f *Food) GetPrice() float64 {
	return f.price
}

func (f *Food) SetPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Food) GetName() string {
	return f.name
}

func (f *Food) SetName(newName string) {
	f.name = newName
}

func (f *Food) ApplyDiscount(discount float64) {
	f.price = (100 - discount) * f.price / 100
}

func (f *Food) ProductInfo() string {
	return fmt.Sprintf("Название: %s, Вес: %.2f kg, Цена: %.2f", f.name, f.weight, f.price)
}
