package lab7

import "fmt"

type Vegetables struct {
	Name   string
	Weight float64
	Price  float64
}

func (v *Vegetables) GetName() string {
	return v.Name
}

func (v *Vegetables) GetPrice() float64 {
	return v.Price
}

func (v *Vegetables) SetPrice(price float64) {
	v.Price = price
}

func (v *Vegetables) ApplyDiscount(discount float64) {
	v.Price -= v.Price * discount / 100
}
