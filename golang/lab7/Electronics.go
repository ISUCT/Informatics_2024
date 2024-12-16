package lab7

import "fmt"

type Electronics struct {
	Name  string
	Model string
	Price float64
}

func NewElectronics(name string, model string, price float64) *Electronics {
	e := new(Electronics)
	e.Name = name
	e.Model = model
	e.Price = price
	return e
}

func (e *Electronics) SetPrice(price float64) { e.Price = price }
func (e Electronics) GetPrice() float64       { return e.Price }
func (e *Electronics) SetName(name string)    { e.Name = name }
func (e Electronics) GetName() string         { return e.Name }
func (e *Electronics) SetModel(model string)  { e.Model = model }

func (e *Electronics) ApplyDiscount(discount float64) {
	if discount > 0 && discount <= 100 {
		e.SetPrice(e.GetPrice() * (1 - discount/100))
	} else {
		fmt.Println("Неверная скидка")
	}
}
