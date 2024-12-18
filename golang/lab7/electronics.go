package lab7

type Electronics struct {
	price  float64
	colour string
	memory string
}

func (x Electronics) GetPrice() float64 {
	return x.price
}

func (c *Electronics) ApplyingTheDiscount(x float64) {
	(*c).price = (c.price / 100) * (100 - x)
}

func (c *Electronics) ChangePrice(x float64) {
	(*c).price = x
}

func (c *Electronics) ChangeCharacteristic(newColour string, newMemory string) {
	c.colour = newColour
	c.memory = newMemory
}
