package labs

import "fmt"

type Cosmetics struct {
	name  string
	price float64
	brand string
}

func (c *Cosmetics) getName() string {
	return c.name
}
func (c *Cosmetics) getPrice() float64 {
	return c.price
}

func (c *Cosmetics) setdiscount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Cosmetics) setPrice(newPrice float64) {
	c.price = newPrice
}
func (c *Cosmetics) GetInfo() string {
	return fmt.Sprintf("Название: %s, Бренд: %s, Цена: %.2f", c.name, c.brand, c.price)
}
