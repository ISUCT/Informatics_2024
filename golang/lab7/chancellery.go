package lab7

import "fmt"

type Chancellery struct {
	Name  string
	Viev  string
	Price float64
}

func (c Chancellery) GetInfo() {
	fmt.Println("В наличии есть", c.Name, "которая", c.Viev, "и имеет цену", c.Price)
}

func (c Chancellery) GetPrice() float64 {
	return c.Price
}

func (c *Chancellery) ApplyDiscount(x float64) {
	(*c).Price = (c.Price / 100) * (100 - x)
}

func (c *Chancellery) EditPrice(x float64) {
	(*c).Price = x
}

func (c *Chancellery) EditDescription(x string) {
	(*c).Viev = x
}
