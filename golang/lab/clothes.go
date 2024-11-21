package lab

import "fmt"

type Clothes struct {
	Name  string
	Size  string
	Price float32
}

func (c Clothes) GetInfo() {
	fmt.Println("У нас есть в наличии", c.Name, "размера", c.Size, "стоимостью", c.Price)
}

func (c Clothes) GetPrice() float32 {
	return c.Price
}

func (c *Clothes) MakeDiscount(x float32) {
	(*c).Price = (c.Price / 100) * (100 - x)
}

func (c *Clothes) ChangePrice(x float32) {
	(*c).Price = x
}

func (c *Clothes) ChangeDescription(x string) {
	(*c).Size = x
}
