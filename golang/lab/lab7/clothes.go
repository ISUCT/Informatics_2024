package lab7

import "fmt"

type Clothes struct {
	Name  string
	Size  string
	Price float64
}

func (c Clothes) GetInfo() {
	fmt.Println("У нас есть в наличии", c.Name, "размера", c.Size, "стоимостью", c.Price)
}

func (c Clothes) GetPrice() float64 {
	return c.Price
}

func (c *Clothes) MakeDiscount(x float64) {
	(*c).Price = (c.Price / 100) * (100 - x)
}

func (c *Clothes) ChangePrice(x float64) {
	(*c).Price = x
}

func (c *Clothes) ChangeDescription(x string) {
	(*c).Size = x
}
