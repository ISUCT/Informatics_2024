package lab7

import (
	"fmt"
)

type Food struct {
	Name   string
	Price  float64
	Weight float64
}

func (f *Food) applyDiscount(discount float64) error {

	if err := validateDiscount(discount); err != nil {
		return err
	}
	f.Price = f.Price * (1 - discount/100)
	return nil
}

func (f *Food) getPrice() float64 {
	return f.Price
}

func (f *Food) getProductInfo() string {
	if f.Name == "" {
		return ("Не все поля заполнены!!!")
	}
	return fmt.Sprintf("Name: %s, Weight: %.2f kg, Price: %.2f", f.Name, f.Weight, f.Price)
}
