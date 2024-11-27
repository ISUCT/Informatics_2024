package lab6

import "fmt"

type City struct {
	Name  string
	Area  string
	Popul string
}

func NewCity(name, area, popul string) *City {
	c := new(City)
	c.Name = name
	c.Area = area
	c.Popul = popul
	return c
}

func (c *City) UpdateStruct(new_name, new_area, new_popul string) {
	c.Name = new_name
	c.Area = new_area
	c.Popul = new_popul
}

func (c City) GetName() string  { return c.Name }
func (c City) GetArea() string  { return c.Area }
func (c City) GetPopul() string { return c.Popul }

func RunLab6Tasks() {
	city := NewCity("Москва", "Россия", "12.6 млн")
	fmt.Println("Население", city.GetPopul())
	fmt.Println("Страна", city.GetArea())
}
