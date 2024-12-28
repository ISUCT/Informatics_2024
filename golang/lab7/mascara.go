package lab7

import "fmt"

type Mascara struct {
	name  string
	price float64
	brand string
}

func (m *Mascara) getName() string {
	return m.name
}
func (m *Mascara) getPrice() float64 {
	return m.price
}

func (m *Mascara) setdiscount(discount float64) {
	m.price -= m.price * discount / 100
}

func (m *Mascara) setPrice(newPrice float64) {
	m.price = newPrice
}
func (m *Mascara) GetInfo() string {
	return fmt.Sprintf("Продается товар: %s, бренда: %s, ценой: %.2f", m.name, m.brand, m.price)
}
