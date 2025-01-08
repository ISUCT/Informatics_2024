package lab7

import "fmt"

type CatFood struct {
	Name      string
	Brand     string
	Price     float64
}

func (f CatFood) GetDescription() {
	fmt.Println("Корм для кошек", f.Name, "от бренда", f.Brand, "он стоит", f.Price)
}

func (f CatFood) GetPrice() float64 {
	return f.Price
}

func (f *CatFood) ApplyDiscount(x float64) {
	(*f).Price = (f.Price) * (x)
}

func (f *CatFood) UpdatePrice(x float64) {
	(*f).Price = x
}

func (f *CatFood) UpdateDescription(x string) {
	(*f).Name = x
}
