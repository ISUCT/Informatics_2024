package labs

import "fmt"

type Cosmetics struct {
	name  string
	price float64
	brand string
}

func (c *Cosmetics) GetName() string {
	return c.name
}
func (c *Cosmetics) GetPrice() float64 {
	return c.price
}

func (c *Cosmetics) Setdiscount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Cosmetics) SetPrice(newPrice float64) {
	c.price = newPrice
}
func (c *Cosmetics) GetInfo() string {
	return fmt.Sprintf("Название: %s, Бренд: %s, Цена: %.2f", c.name, c.brand, c.price)
}
