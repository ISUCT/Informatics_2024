package lab6

import (
	"fmt"
)

// ship
type Ship struct {
	Name         string
	Lenght       int
	Displacement int
}

func NewShip(name string, lenght int, displacement int) Ship {
	return Ship{
		Name:         name,
		Lenght:       lenght,
		Displacement: displacement,
	}
}
func (f *Ship) GetLenght() int {
	return f.Lenght
}
func (f *Ship) SetLenght(lenght int) {
	f.Lenght = lenght
}
func (f *Ship) GetDisplacement() int {
	return f.Displacement
}
func (f *Ship) SetDisplacement(displacement int) {
	f.Displacement = displacement
}
func Lab6() {
	ship1 := NewShip("Tight", 150, 3000)
	fmt.Println("Ship Lenght:", ship1.GetLenght())
	ship1.SetLenght(250)
	fmt.Println("New Lenght:", ship1.GetLenght())
	fmt.Println("Ship Displacement:", ship1.GetDisplacement())
	ship1.SetDisplacement(3500)
	fmt.Println("New Displacement:", ship1.GetDisplacement())
}
