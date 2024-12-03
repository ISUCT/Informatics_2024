package lab7

type Skateboards struct {
	price float32
	color string
	material string
}

func (b Skateboards) GetPrice() float32 {
	return b.price
}

func (b *Skateboards) Discount(x float32) {
	(*b).price = (b.price / 100) * (100 - x)
}

func (b *Skateboards) ChangePrice(x float32) {
	(*b).price = x
}

func (f *Skateboards) ChangeCharacteristic(newcolor string, newmaterial string) {
	f.color = newcolor
	f.material = newmaterial
}