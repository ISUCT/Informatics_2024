package components

import (
	"errors"
	"fmt"
)

type Electronics struct {
	Name  string
	Price float64
	Brand string
	Model string
}

func (e *Electronics) sale(discount float64) error {
	if e.Name == "" || e.Brand == "" || e.Model == "" {
		return errors.New("Не все поля заполнены для Electronics")
	}
	if discount < 0 {
		return errors.New("Скидка не может быть меньше 0")
	}
	if discount > 100 {
		return errors.New("Скидка не может больше 100%")
	}
	e.Price = e.Price * (1 - discount/100)
	return nil
}

func (e *Electronics) price() float64 {
	return e.Price
}

func (e *Electronics) productFeature() string {
	if e.Name == "" || e.Brand == "" || e.Model == "" {
		return ("Не все поля заполнены!!!")
	}
	return fmt.Sprintf("Name: %s, Brand: %s, Model: %s, Price: %.2f", e.Name, e.Brand, e.Model, e.Price)
}
