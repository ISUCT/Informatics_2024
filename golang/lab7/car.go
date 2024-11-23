package lab7

type Car struct {
	price  float32
	colour string
	model  string
}

func (x Car) GetPrice() float32 {
	return x.price
}

func (c *Car) Sale(x float32) {
	(*c).price = (c.price / 100) * (100 - x)
}

func (c *Car) ChangePrice(x float32) {
	(*c).price = x
}

func (c *Car) changeCharacteristic(newcolour string, newmodel string) {
	c.colour = newcolour
	c.model = newmodel
}
