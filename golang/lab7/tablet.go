package lab7

type Tablet struct {
	Name  string
	Price float64
}

func (t *Tablet) GetName() string {
	return t.Name
}

func (t *Tablet) GetPrice() float64 {
	return t.Price
}

func (t *Tablet) ApplyDiscount(discount float64) {
	t.Price -= t.Price * discount / 100
}

