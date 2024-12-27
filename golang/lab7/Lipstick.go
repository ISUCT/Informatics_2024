package lab7

import (
	"fmt"
)

type Lipstick struct {
	Brand   string
	Color string
	Price  float64
}

func (l Lipstick) GetInfo() {
	fmt.Println("В магазине продаётся", l.Color, l.Brand, "ценой", l.Price, "руб.")
}

func (l Lipstick) GetPrice() float64 {
	return l.Price
}

func (l *Lipstick) SetDiscount(x float64) {
	(*l).Price = (l.Price / 100) * (100 - x)
}

func (l *Lipstick) EditPrice(x float64) {
	(*l).Price = x
}

func (l *Lipstick) EditDescription(x string) {
	(*l).Color = x
}
