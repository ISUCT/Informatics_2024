package lab7

import "fmt"

type Magazine struct {
	Name      string
	Publisher string
	Price     float64
}

func NewMagazine(name string, publisher string, price float64) *Magazine {
	m := new(Magazine)
	m.Name = name
	m.Publisher = publisher
	m.Price = price
	return m
}

func (m *Magazine) SetPrice(price float64)        { m.Price = price }
func (m Magazine) GetPrice() float64              { return m.Price }
func (m *Magazine) SetName(name string)           { m.Name = name }
func (m Magazine) GetName() string                { return m.Name }
func (m *Magazine) SetPublisher(publisher string) { m.Publisher = publisher }

func (m *Magazine) ApplyDiscount(discount float64) {
	if discount > 0 && discount <= 100 {
		m.SetPrice(m.GetPrice() * (1 - discount/100))
	} else {
		fmt.Println("Неверная скидка")
	}
}
