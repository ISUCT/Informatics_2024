package lab7

import "fmt"

type Car struct {
	Name  string
	Model string
	Price float64
}

func NewCar(name string, model string, price float64) *Car {
	e := new(Car)
	e.Name = name
	e.Model = model
	e.Price = price
	return e
}

func (e *Car) SetPrice(price float64) { e.Price = price }
func (e Car) GetPrice() float64       { return e.Price }
func (e *Car) SetName(name string)    { e.Name = name }
func (e Car) GetName() string         { return e.Name }
func (e *Car) SetModel(model string)  { e.Model = model }

func (e *Car) ApplyDiscount(discount float64) {
	if discount > 0 && discount <= 100 {
		e.SetPrice(e.GetPrice() * (1 - discount/100))
	} else {
		fmt.Println("Неверная скидка")
	}
}
