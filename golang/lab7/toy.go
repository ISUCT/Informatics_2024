package lab7

type Toy struct {
	name     string
	price    float64
	material string
}

func (t *Toy) GetName() string {
	return t.name
}

func (t *Toy) SetName(NewName string) {
	t.name = NewName
}

func (t *Toy) GetPrice() float64 {
	return t.price
}

func (t *Toy) SetPrice(NewPrice float64) {
	t.price = NewPrice
}

func (t *Toy) ApplyDiscount(discount float64) {
	t.price = t.price * (100 - discount) / 100
}

func (t *Toy) SetMaterial(NewMaterial string) {
	t.material = NewMaterial
}
