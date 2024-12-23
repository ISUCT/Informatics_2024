package lab7

import (
	"fmt"
)

type Food struct {
	Price  float64
	Name   string
	Weight float64
}

func (f *Food) GetPrice() float64 {
	return f.Price
}

func (f *Food) SetPrice(newPrice float64) {
	if newPrice >= 0 {
		f.Price = newPrice
	}
}

func (f *Food) GetName() string {
	return f.Name
}

func (f *Food) SetName(newName string) {
	f.Name = newName
}

func (f *Food) ApplyDiscount(discount float64) {
	f.Price = (100 - discount) * f.Price / 100
}

func (f *Food) ProductInfo() string {
	return fmt.Sprintf("Название продукта: %s, Вес: %.2f kg, Цена: %.2f", f.Name, f.Weight, f.Price)
}
