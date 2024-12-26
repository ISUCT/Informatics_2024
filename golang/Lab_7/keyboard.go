package lab7

import "fmt"

type Keyboard struct {
	Brand      string
	Format     string
	NumberKeys int
	Price      float64
}

func (k Keyboard) GetPrice() float64 {
	return k.Price
}

func (k *Keyboard) EditPrice(NewPrice float64) {
	(*k).Price = NewPrice
}

func (k *Keyboard) ApplyDiscount(Percent float64) {
	(*k).Price = (k.Price / 100) * (100 - Percent)
}

func (k *Keyboard) EditDescription(NewDescription string) {
	(*k).Format = NewDescription
}

func (k Keyboard) GetInfo() {
	fmt.Println("В продаже есть клавиатура, ", k.Format, "бренда", k.Brand, ", у которой", k.NumberKeys, " клавиш и стоит:", k.Price)
}
