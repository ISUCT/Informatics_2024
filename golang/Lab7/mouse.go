package lab7

import "fmt"

type Mouse struct {
	Brand  string
	Sensor string
	Price  float64
}

func (m Mouse) GetPrice() float64 {
	return m.Price
}

func (m *Mouse) EditPrice(NewPrice float64) {
	(*m).Price = NewPrice
}

func (m *Mouse) ApplyDiscount(Percent float64) {
	(*m).Price = (m.Price / 100) * (100 - Percent)
}

func (m *Mouse) EditDescription(NewDescription string) {
	(*m).Sensor = NewDescription
}

func (m Mouse) GetInfo() {
	fmt.Println("В продаже есть компьютерная мышь от производителя", m.Brand, "с ", m.Sensor, "сенсором и стоимостью:", m.Price)
}
