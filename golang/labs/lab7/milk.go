package lab7

import "fmt"

type Milk struct {
	Name    string
	Price   float64
	Percent float64
}

func NewMilk(name string, price float64, percent float64) *Milk {
	milk := &Milk{Name: name, Price: price, Percent: percent}
	return milk
}

func (m *Milk) GetName() string {
	return m.Name
}

func (m *Milk) GetPrice() float64 {
	return m.Price
}

func (m *Milk) SetPrice(price float64) {
	m.Price = price
}

func (m *Milk) GetPercent() float64 {
	return m.Percent
}

func (m *Milk) ApplyDiscount(discount float64) {
	m.Price -= m.Price * discount / 100
}

func (m *Milk) GetInfo() {
	fmt.Printf("name: %v\nprice: %.2f\npercent: %.2f%%\n", m.Name, m.Price, m.Percent)
}
