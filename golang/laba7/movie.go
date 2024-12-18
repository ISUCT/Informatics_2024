package laba7

import "fmt"

type Movie struct {
	Price    float64
	Title    string
	Director string
	Duration int
}

func (m *Movie) GetPrice() float64 {
	return m.Price
}

func (m *Movie) ApplyDiscount(discount float64) {
	m.Price = m.Price * (1 - discount/100)
}

func (m *Movie) GetProductInfo() string {
	return fmt.Sprintf("Title: %s, Director: %s, Duration: %d, Price: %.2f", m.Title, m.Director, m.Duration, m.Price)
}
