package lab7

type Car struct {
	Name  string
	Price float64
	Color string
}

func (c *Car) GetName() string {
	return c.Name
}

func (c *Car) GetPrice() float64 {
	return c.Price
}

func (c *Car) SetPrice(price float64) {
	c.Price = price
}

func (c *Car) ApplyDiscount(discount float64) {
	c.Price -= c.Price * discount / 100
}
func (c *Car) ChangeColor(newColor string) {
	c.Color = newColor
}