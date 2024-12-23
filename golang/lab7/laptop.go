package lab7

type Laptop struct {
	Name  string
	Price float64
}

func (l *Laptop) GetName() string {
	return l.Name
}

func (l *Laptop) GetPrice() float64 {
	return l.Price
}

func (l *Laptop) ApplyDiscount(discount float64) {
	l.Price -= l.Price * discount / 100
}

