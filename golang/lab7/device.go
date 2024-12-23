package lab7

import (
	"fmt"
)

type Device struct {
	Price float64
	Name  string
	Color string
	Model string
	Brand string
}

func (t *Device) GetPrice() float64 {
	return t.Price
}

func (t *Device) SetPrice(newPrice float64) {
	if newPrice >= 0 {
		t.Price = newPrice
	}
}

func (t *Device) GetName() string {
	return t.Name
}

func (t *Device) GetColor() string {
	return t.Color
}

func (t *Device) GetModel() string {
	return t.Model
}

func (t *Device) GetBrand() string {
	return t.Brand
}

func (t *Device) ApplyDiscount(discount float64) {
	t.Price = (100 - discount) * t.Price / 100
}

func (t *Device) ProductInfo() string {
	return fmt.Sprintf("Устройство: %s, Бренд: %s, Модель: %s, Цвет: %s, Цена: %.2f", t.Name, t.Brand, t.Model, t.Color, t.Price)
}
