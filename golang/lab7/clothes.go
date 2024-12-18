package lab7

type Clothes struct {
	price  float64
	colour string
	size   string
}

func (x Clothes) GetPrice() float64 {
	return x.price
}

func (c *Clothes) ApplyingTheDiscount(x float64) {
	(*c).price = (c.price / 100) * (100 - x)
}

func (c *Clothes) ChangePrice(x float64) {
	(*c).price = x
}

func (c *Clothes) ChangeCharacteristic(newColour string, newSize string) {
	c.colour = newColour
	c.size = newSize
}
