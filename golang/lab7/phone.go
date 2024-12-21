package lab7

type Phone struct {
	price  float64
	name   string
	weight float64
}

func (t *Phone) GetName() string {
	return t.name
}

func (t *Phone) SetName(newName string) {
	t.name = newName
}

func (t *Phone) GetPrice() float64 {
	return t.price
}

func (t *Phone) SetPrice(newPrice float64) {
	t.price = newPrice
}

func (t *Phone) SetWeight(newWeight float64) {
	t.weight = newWeight
}

func (t *Phone) ApplyDiscount(discount float64) {
	t.price = t.price * (100 - discount) / 100
}
