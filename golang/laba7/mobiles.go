package laba7

import "fmt"

type Mobiles struct {
	Name      string
	Color     string
	Price     float64
	SailStart float64
}

func (m *Mobiles) ApplyDiscount(discount float64) {
	m.Price = m.Price * (1 - discount/100)
}

func (m *Mobiles) GetPrice() float64 {
	return m.Price
}
func (m *Mobiles) SetPrice(newPrice float64) {
	m.Price = newPrice
}
func (m *Mobiles) ChangeColor(newColor string) {
	m.Color = newColor
}

func (m *Mobiles) getInfo() string {
	return fmt.Sprintf("Name: %s, Color: %s, Price: %.2f,SailStart: %.2f", m.Name, m.Color, m.Price, m.SailStart)
}
