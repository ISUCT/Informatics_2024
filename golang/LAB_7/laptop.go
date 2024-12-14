package LAB_7

type Laptop struct {
	Name  string
	Price float64
	Brand string
	Color string
}

func (l *Laptop) GetName() string {
	return l.Name
}

func (l *Laptop) GetPrice() float64 {
	return l.Price
}

func (l *Laptop) SetPrice(price float64) {
	l.Price = price
}

func (l *Laptop) ApplyDiscount(discount float64) {
	l.Price -= l.Price * discount / 100
}
func (l *Laptop) ChangeColor(newColor string) {
	l.Color = newColor
}
