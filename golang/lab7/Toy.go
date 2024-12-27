package lab7

import "fmt"

type Toy struct {
	Hero  string
	Size  string
	Price float64
}

func (t Toy) GetInfo() {
	fmt.Println("В магазине продаётся", t.Hero, "размером", t.Size, ", ценой", t.Price, "руб.")
}

func (t Toy) GetPrice() float64 {
	return t.Price
}

func (t *Toy) SetDiscount(x float64) {
	(*t).Price = (t.Price / 100) * (100 - x)
}

func (t *Toy) EditPrice(x float64) {
	(*t).Price = x
}

func (t *Toy) EditDescription(x string) {
	(*t).Size = x
}
