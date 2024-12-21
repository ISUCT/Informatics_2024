package lab7

import "fmt"

type Book struct {
	Name   string
	Format string
	Price  float64
}

func (b *Book) GetInfo() {
	fmt.Println("Книга:", b.Name, "в", b.Format, "формате, стоимость:", b.Price)
}

func (b *Book) GetPrice() float64 {
	return b.Price
}

func (b *Book) MakeDiscount(discountPercentage float64) {
	if discountPercentage >= 100 {
		b.Price = 0
		return
	}
	b.Price = b.Price * (1 - discountPercentage/100)
}

func (b *Book) ChangePrice(newPrice float64) {
	b.Price = newPrice
}

func (b *Book) ChangeDescription(newFormat string) {
	b.Format = newFormat
}
