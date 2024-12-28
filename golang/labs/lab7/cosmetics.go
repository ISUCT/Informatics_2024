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

func (c *Cosmetics) SetDiscount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Cosmetics) SetPrice(newPrice float64) {
	c.price = newPrice
}
func (c *Cosmetics) GetInf() string {
	return fmt.Sprintf("название: %s, бренд: %s, цена: %.2f", c.name, c.brand, c.price)
}
