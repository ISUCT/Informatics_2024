package lab7

import "fmt"

type Electronics struct {
	name  string
	price float64
	brand string
	color string
}

func (e *Electronics) getPrice() float64 {
	return e.price
}

func (e *Electronics) setPrice(newPrice float64) {
	e.price = newPrice
}

func (e *Electronics) applyDiscount(discount float64) {
	if discount < 0 || discount > 100 {
		fmt.Println("Скидка должна быть в диапазоне от 0 до 100%")
		return
	}
	e.price = e.price * (1 - discount/100)
}

func (e *Electronics) getName() string {
	return e.name
}

func (e *Electronics) setName(newName string) {
	e.name = newName
}
