package lab7

import "fmt"

type Car struct {
	Mark     string
	Colour   string
	MaxSpeed int
	Price    float64
}

func (c Car) GetPrice() float64 {
	return c.Price
}

func (c *Car) EditPrice(NewPrice float64) {
	(*c).Price = NewPrice
}

func (c *Car) ApplyDiscount(Percent float64) {
	(*c).Price = (c.Price / 100) * (100 - Percent)
}

func (c *Car) EditDescription(NewDescription string) {
	(*c).Colour = NewDescription
}

func (c Car) GetInfo() {
	fmt.Println("Продаётся машина", c.Mark, "цвета", c.Colour, ", которая разгоняется до", c.MaxSpeed, "и стоит в долларах:", c.Price)
}
