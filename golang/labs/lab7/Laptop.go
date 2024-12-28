package lab7

type Laptop struct {
	Model string
	Price float64
	ROM   float64
}

func (l *Laptop) SetModel(Model string) {
	l.Model = Model
}
func (l *Laptop) GetModel() string {
	return l.Model
}
func (l *Laptop) SetPrice(Price float64) {
	l.Price = Price
}
func (l *Laptop) GetPrice() float64 {
	return l.Price
}

func (l *Laptop) SetROM(ROM float64) {
	l.ROM = ROM
}

func (l *Laptop) SetDiscount(percent float64) {
	l.Price -= (l.Price / 100) * percent
}
