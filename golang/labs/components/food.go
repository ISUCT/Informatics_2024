package components

import (
	"errors"
	"fmt"
)

type Food struct {
	Name   string
	Price  float64
	Weight float64
}

func (f *Food) sale(discount float64) error {
	if f.Name == "" {
		return errors.New("Не все поля заполнены для Food")
	}
	if discount < 0 {
		return errors.New("Скидка не может быть меньше 0")
	}
	if discount > 100 {
		return errors.New("Скидка не может больше 100%")
	}
	f.Price = f.Price * (1 - discount/100)
	return nil
}

func (f *Food) price() float64 {
	return f.Price
}

func (f *Food) productFeature() string {
	if f.Name == "" {
		return ("Не все поля заполнены!!!")
	}
	return fmt.Sprintf("Name: %s, Weight: %.2f kg, Price: %.2f", f.Name, f.Weight, f.Price)
}
