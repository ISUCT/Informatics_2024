package lab7

type PC struct {
	Model string
	Price float64
	RAM   float64
}

func (c *PC) SetModel(Model string) {
	c.Model = Model
}

func (c *PC) SetPrice(Price float64) {
	c.Price = Price
}

func (c *PC) GetPrice() float64 {
	return c.Price
}

func (c *PC) GetModel() string {
	return c.Model
}

func (c *PC) SetDiscount(percent float64) {
	c.Price -= (c.Price / 100) * percent
}

func (c *PC) SetRAM(RAM float64) {
	c.RAM = RAM
}

func (c *PC) GetRAM() float64 {
	return c.RAM
}
