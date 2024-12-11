package laba7

import "fmt"

type Movie struct {
	price    float64
	title    string
	director string
	duration int
}

func (m *Movie) getPrice() float64 {
	return m.price
}

func (m *Movie) applyDiscount(discount float64) {
	m.price = m.price * (1 - discount/100)
}

func (m *Movie) getProductInfo() string {
	return fmt.Sprintf("Title: %s, Director: %s, Duration: %d, Price: %.2f", m.title, m.director, m.duration, m.price)
}
