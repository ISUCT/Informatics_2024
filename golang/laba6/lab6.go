package lab6

import (
	"fmt"
)

type Country struct {
	Name       string
	Capital    string
	Area       float64
	Population int
}

func NewCountry(name, capital string, area float64, population int) *Country {
	return &Country{
		Name:       name,
		Capital:    capital,
		Area:       area,
		Population: population,
	}
}

func (p Country) GetCapital() string {
	return p.Capital
}
func (p *Country) SetArea(area float64) {
	p.Area = area
}
func (p Country) GetArea() float64 {
	return p.Area
}
func (p Country) GetPopulation() int {
	return p.Population
}

func Laba6Run() {
	France := NewCountry("Франция", "Париж", 550.695, 67000000)
	France.SetArea(551.695)

	fmt.Printf("country: %sn", France.Name)
	fmt.Printf("capital: %sn", France.GetCapital())
	fmt.Printf("area: %.2f км²n", France.GetArea())
	fmt.Printf("population: %dn", France.GetPopulation())
}