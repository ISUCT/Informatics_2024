package lab7

type eyeliner struct {
	price float32
	color string
	material string
}

func (x eyeliner) GetPrice() float32 {
	return x.price
}

func (c *eyeliner) Discount(x float32) {
 (*c).price = (c.price / 100) * (100 - x)
}

func (c *eyeliner ChangePrice(x float32) {
 (*c).price = x
}

func (c *eyeliner) ChangeCharacteristic(newcolor string,newmaterial string ) {
	c.color = newcolor
	c.material = newmaterial
}
