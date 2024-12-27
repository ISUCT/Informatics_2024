package lab7

import (
	"fmt"
)

type Drink struct {
	Name   string
	Volume string
	Price  float64
}

func (d Drink) GetInfo() {
	fmt.Println("В магазине продаётся", d.Name, "объёмом", d.Volume, ", ценой", d.Price, "руб.")
}

func (d Drink) GetPrice() float64 {
	return d.Price
}

func (d *Drink) SetDiscount(x float64) {
	(*d).Price = (d.Price / 100) * (100 - x)
}

func (d *Drink) EditPrice(x float64) {
	(*d).Price = x
}

func (d *Drink) EditDescription(x string) {
	(*d).Volume = x
}
