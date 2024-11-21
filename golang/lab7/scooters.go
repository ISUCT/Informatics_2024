package lab7

type Scooters struct {
	price float32
	color string
	material string
}

func (b Scooters) GetPrice() float32 {
	return b.price
}

func (b *Scooters) Sale(x float32) {
	(*b).price = (b.price / 100) * (100 - x)
}

func (b *Scooters) ChangePrice(x float32) {
	(*b).price = x
}

func (b *Scooters) changeCharacteristic(newcolor string, newmaterial string) {
	b.color = newcolor
	b.material = newmaterial
}