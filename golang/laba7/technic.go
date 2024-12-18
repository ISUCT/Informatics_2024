package laba7

import (
	"fmt"
)

type Technic struct {
	brand string
	price float64
	model string
	color string
	name  string
	code  float64
}

func (t *Technic) ApplyDiscount(discount float64) {
	t.price = (100 - discount) * t.price / 100
}

func (t *Technic) GetPrice() float64 {
	return t.price
}

func (t *Technic) SetPrice(newPrice float64) {
	t.price = newPrice
}

func (t *Technic) GetName() string {
	return t.name
}

func (t *Technic) GetColor() string {
	return t.color
}

func (t *Technic) GetModel() string {
	return t.model
}

func (t *Technic) GetCode() float64 {
	return t.code
}

func (t *Technic) GetBrand() string {
	return t.brand
}

func (t *Technic) ProductInfo() string {
	return fmt.Sprintf("Техника: %s, Бренд: %s, Модель техники: %s, Цвет товара: %s,Код товара: %.2f, Цена товара: %.2f",
		t.name, t.brand, t.model, t.color, t.code, t.price)
}
