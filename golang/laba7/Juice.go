package laba7

import "fmt"

type Juice struct {
	Name   string
	Volume float64
	Taste  string
	Price  float64
}

func (j *Juice) ApplyDiscount(discount float64) {
	j.Price = j.Price * (1 - discount/100)
}

func (j *Juice) GetPrice() float64 {
	return j.Price
}

func (j *Juice) GetProductInfo() string {
	return fmt.Sprintf("Name: %s, Taste: %s, Volume: %.2f, Price: %.2f", j.Name, j.Taste, j.Volume, j.Price)
}
