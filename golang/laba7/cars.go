package laba7

import (
	"fmt"
)

type Car struct {
	price float64
	name  string
	color string
	model string
	brand string
}

func (c *Car) GetPrice() float64 {
	return c.price
}

func (c *Car) SetPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Car) ApplyDiscount(discount float64) {
	c.price = (100 - discount) * c.price / 100
}

func (c *Car) GetModel() string {
	return c.model
}

func (c *Car) GetName() string {
	return c.name
}

func (c *Car) GetBrand() string {
	return c.brand
}

func (c *Car) GetColor() string {
	return c.color
}

func (c *Car) SetColor(newName string) {
	c.color = newName
}

func (c *Car) ProductInfo() string {
	return fmt.Sprintf("Машина: %s, Бренд: %s, Модель машины: %s, Цвет машины: %s, Цена машины: %.2f", c.name, c.brand, c.model, c.color, c.price)
}
