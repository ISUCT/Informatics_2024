package main

import (
	"fmt"
)

type Electronics struct {
	name  string
	brand string
	price float64
}

func (e *Electronics) GetPrice() float64 {
	return e.price
}

func (e *Electronics) SetPrice(price float64) {
	e.price = price
}

func (e *Electronics) ApplyDiscount(percent float64) {
	e.price -= e.price * percent / 100
}
func (e *Electronics) GetDetails() string {
	return fmt.Sprintf("Electronics: %s, Brand: %s, Price: %.2f", e.name, e.brand, e.price)
}
