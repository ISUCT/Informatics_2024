package lab7

type Electronics struct {
	price  float32
	colour string
	memory string
}

func (x Electronics) GetPrice() float32 {
	return x.price
}

func (c *Electronics) Sale(x float32) {
	(*c).price = (c.price / 100) * (100 - x)
}

func (c *Electronics) ChangePrice(x float32) {
	(*c).price = x
}

func (c *Electronics) changeCharacteristic(newcolour string, newmemory string) {
	c.colour = newcolour
	c.memory = newmemory
}
