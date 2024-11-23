package lab7

type Clothes struct {
	price  float32
	colour string
	size   string
}

func (x Clothes) GetPrice() float32 {
	return x.price
}

func (c *Clothes) Sale(x float32) {
	(*c).price = (c.price / 100) * (100 - x)
}

func (c *Clothes) ChangePrice(x float32) {
	(*c).price = x
}

func (c *Clothes) changeCharacteristic(newcolour string, newsize string) {
	c.colour = newcolour
	c.size = newsize
}
