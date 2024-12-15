package lab7

type Car struct {
	price  float64
	colour string
	model  string
}

func (x Car) GetPrice() float64 {
	return x.price
}

func (c *Car) ApplyingTheDiscount(x float64) {
	(*c).price = (c.price / 100) * (100 - x)
}

func (c *Car) ChangePrice(x float64) {
	(*c).price = x
}

func (c *Car) ChangeCharacteristic(newcolour string, newmodel string) {
	c.colour = newcolour
	c.model = newmodel
}
