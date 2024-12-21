package Lab7

import (
	"fmt"
)

type Device struct {
	Name  string
	Price float64
	Brand string
	Model string
}

func (e *Device) applyDiscount(discount float64) error {
	if err := validateDiscount(discount); err != nil {
		return err
	}
	e.Price = e.Price * (1 - discount/100)
	return nil
}

func (e *Device) getPrice() float64 {
	return e.Price
}

func (e *Device) getProductInfo() string {
	if e.Name == "" || e.Brand == "" || e.Model == "" {
		return "Не все поля заполнены"
	}
	return fmt.Sprintf("Девайс: %s, Бренд: %s, Модель: %s, Цена: %.2f", e.Name, e.Brand, e.Model, e.Price)
}
