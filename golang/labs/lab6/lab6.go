package lab6

import "fmt"

type Plane struct {
	Name        string
	Model       string
	FlightSpeed float64
}

func NewPlane(name, model string, flightspeed float64) Plane {
	return Plane{
		Name:        name,
		Model:       model,
		FlightSpeed: flightspeed,
	}
}
func (p *Plane) SetFlightSpeed(speed float64) {
	p.FlightSpeed = speed
}
func (p *Plane) GetFlightSpeed() float64 {
	return p.FlightSpeed
}
func (p Plane) DisplayInfo() {
	fmt.Printf("Название: %s,Модель: %s,Скорость полёта: %.2f km/h\n", p.Name, p.Model, p.FlightSpeed)
}
func Runlab6() {

	plane := NewPlane("Albatros", "718", 900.0)

	plane.DisplayInfo()

	speed := plane.GetFlightSpeed()
	fmt.Printf("Текущая скорость полёта:%.2f km/h\n", speed)
	plane.SetFlightSpeed(950.0)
	plane.DisplayInfo()
}
