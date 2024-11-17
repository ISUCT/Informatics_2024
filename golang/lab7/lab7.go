package lab7

import (
	"fmt"
)

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}
type Item struct {
	Name     string
	Price    float64
	Category string
}

func (i *Item) GetName() string {
	return i.Name
}
func (i *Item) GetPrice() float64 {
	return i.Price
}
func (i *Item) SetPrice(price float64) {
	i.Price = price
}
func (i *Item) ApplyDiscount(discount float64) {
	i.Price = i.Price * (1 - discount/100)
}
func CalculateTotal(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}
func Lab7() {
	item1 := &Item{Name: "Laptop", Price: 1000, Category: "Electronics"}
	item2 := &Item{Name: "Phone", Price: 500, Category: "Electronics"}
	item3 := &Item{Name: "Book", Price: 20, Category: "Literature"}
	item1.ApplyDiscount(10) // 10% скидка.
	item2.ApplyDiscount(5)  // 5% скидка
	products := []Product{item1, item2, item3}
	// стоимость до применения скидок
	fmt.Println("Общая стоимость товаров:", CalculateTotal(products))
	// Изменяем цену товара
	item3.SetPrice(15) // Изменяем цену книги
	// стоимость после изменения цены
	fmt.Println("Общая стоимость товаров после изменения цены: ", CalculateTotal(products))
}
