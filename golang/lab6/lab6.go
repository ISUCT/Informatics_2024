package lab6

import (
	"fmt"
)

type Ship struct {
	Name         string
	Length       float64
	Displacement float64
	Crew         int
}

func NewShip(name string, length float64, displacement float64, crew int) *Ship {
	return &Ship{
		Name:         name,
		Length:       length,
		Displacement: displacement,
		Crew:         crew,
	}
}

func (p Ship) GetLength() float64 {
	return p.Length
}

func (p Ship) GetDisplacement() float64 {
	return p.Displacement
}

func (p *Ship) SetCrew(crew int) {
	p.Crew = crew
}

func (p Ship) GetCrew() int {
	return p.Crew
}

func CompleteLab6() {
	Ship := NewShip("Ямато", 263.0, 72800, 1930)
	Ship.SetCrew(2400)

	fmt.Printf("ship: %s\n", Ship.Name)
	fmt.Printf("length: %.2f\n", Ship.GetLength())
	fmt.Printf("displacement: %.2f\n", Ship.GetDisplacement())
	fmt.Printf("crew: %d\n", Ship.GetCrew())
}
