package lab7

type Technic struct {
	price  float64
	name   string
	weight float64
}

func (t *Technic) getName() string {
	return t.name
}

func (t *Technic) setName(newName string) {
	t.name = newName
}

func (t *Technic) getPrice() float64 {
	return t.price
}

func (t *Technic) setPrice(newPrice float64) {
	t.price = newPrice
}

func (t *Technic) setWeight(newWeight float64) {
	t.weight = newWeight
}

func (t *Technic) applyDiscount(discount float64) {
	t.price = t.price * (100 - discount) / 100
}
