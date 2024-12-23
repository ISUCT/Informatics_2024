package lab7

import (
	"fmt"
)

type Car struct {
	Price float64
	Name  string
	Color string
	Model string
	Brand string
}

func (c *Car) GetPrice() float64 {
	return c.Price
}

func (c *Car) SetPrice(newPrice float64) {
	if newPrice >= 0 {
		c.Price = newPrice
	}
}

func (c *Car) GetModel() string {
	return c.Model
}

func (c *Car) GetName() string {
	return c.Name
}

func (c *Car) GetBrand() string {
	return c.Brand
}

func (c *Car) GetColor() string {
	return c.Color
}

func (c *Car) SetColor(newName string) {
	c.Color = newName
}

func (c *Car) ApplyDiscount(discount float64) {
	c.Price = (100 - discount) * c.Price / 100
}

func (c *Car) ProductInfo() string {
	return fmt.Sprintf("Машина: %s, Бренд: %s, Модель: %s, Цвет: %s, Цена: %.2f", c.Name, c.Brand, c.Model, c.Color, c.Price)
}
