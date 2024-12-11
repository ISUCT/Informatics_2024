package laba7

import "fmt"

type Juice struct {
	name   string
	volume float64
	taste  string
	price  float64
}

func (j *Juice) applyDiscount(discount float64) {
	j.price = j.price * (1 - discount/100)
}

func (j *Juice) getPrice() float64 {
	return j.price
}

func (j *Juice) getProductInfo() string {
	return fmt.Sprintf("Name: %s, Taste: %s, Volume: %.2f, Price: %.2f", j.name, j.taste, j.volume, j.price)
}
