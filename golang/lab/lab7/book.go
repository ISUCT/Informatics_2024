package lab7

import "fmt"

type Book struct {
	Name   string
	Format string
	Price  float32
}

func (b Book) GetInfo() {
	fmt.Println("У нас есть в наличии", b.Name, "в", b.Format, "формате и стоит", b.Price)
}

func (b Book) GetPrice() float32 {
	return b.Price
}

func (b *Book) MakeDiscount(x float32) {
	(*b).Price = (b.Price / 100) * (100 - x)
}

func (b *Book) ChangePrice(x float32) {
	(*b).Price = x
}

func (b *Book) ChangeDescription(x string) {
	(*b).Format = x
}
