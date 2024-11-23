package labs

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

func (c *Cosmetics) discount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Cosmetics) setPrice(newPrice float64) {
	c.price = newPrice
}
