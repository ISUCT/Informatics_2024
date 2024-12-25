package lab7

import "fmt"

type Electronics struct {
	name string
	model string
	price float64
	description string
}

func NewElectronics(name string, model string, price float64, description string) *Electronics {
	electronics := &Electronics{name: name, model: model, price: price, description: description}
	return electronics
}

func (e *Electronics) GetName() string {
	return e.name
}

func (e *Electronics) SetPrice(newPrice float64) {
	e.price = newPrice
}

func (e *Electronics) GetPrice() float64 {
	return e.price
}

func (e *Electronics) ChangeData(name string, price float64, description string) {
	e.name = name
	e.price = price
	e.description = description
}

func (e *Electronics) GetData() {
	fmt.Printf("name: %s\nmodel: %s\nprice: %.2f\ndescription: %s\n\n", e.name, e.model, e.price, e.description)
}

func (e *Electronics) ApplyDiscount(perDiscount float64) {
	e.price = e.price * (1 - (perDiscount / 100))
}