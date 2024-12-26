package lab7

import (
	"errors"
)

type Furniture struct {
	price float64
	name  string
	brand string
	color string
}

func (f *Furniture) Sale(discount float64) error {
	if (0 > discount) || (100 < discount) {
		return errors.New("Скидка должна быть в диапозоне от 0 до 100 процентов")
	}

	f.price -= discount * f.price
	return nil
}

func (f *Furniture) Change_price(new_price float64) {
	f.price = new_price
}

func (f *Furniture) Change_specifications(new_name, new_brand, new_color string) {
	f.brand = new_brand
	f.name = new_name
	f.color = new_color
}

func (f *Furniture) Get_price() float64 {
	return f.price
}

func (f *Furniture) Get_info() [3]string {
	arr := [3]string{f.name, f.brand, f.color}
	return arr
}
