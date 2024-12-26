package lab7

import "fmt"

type Laptop struct {
	Name  string
	GB    float64
	Price float64
	CPU   string
}

func (l *Laptop) ApplyDiscount(discount float64) {
	l.Price = l.Price * (1 - discount/100)
}

func (l *Laptop) GetPrice() float64 {
	return l.Price
}
func (l *Laptop) SetPrice(newPrice float64) {
	l.Price = newPrice
}

func (l *Laptop) ChangeCPU(newCPU string) {
	l.CPU = newCPU
}

func (l *Laptop) getInfo() string {
	return fmt.Sprintf("Name: %s, GB: %.2f , Price: %.2f,CPU: %s", l.Name, l.GB, l.Price, l.CPU)
}
