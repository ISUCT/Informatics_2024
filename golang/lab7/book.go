package lab7

import "fmt"

type Book struct {
	Name  string
	Format string
	Price float64
}

func (b *Book) GetInfo() {
	fmt.Println("Книга:", b.Name, "в", b.Format, "формате, стоимость:", b.Price)
}

func (b *Book) GetPrice() float64 {
	return b.Price
}

func (b *Book) MakeDiscount(x float64) {
	if x >= 100 {
		b.Price = 0
		return
	}
	b.Price = b.Price * (1 - x/100)
}

func (b *Book) ChangePrice(x float64) {
	b.Price = x
}

func (b *Book) ChangeDescription(x string) {
	b.Format = x
}