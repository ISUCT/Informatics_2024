package lab7

import "fmt"

// Интерфейс Product
type Product interface {
	GetPrice() float64
	SetPrice(price float64)
	GetDiscount() float64
	SetDiscount(discount float64)
	GetName() string
	SetName(name string)
}

// Структура ProductBase
type ProductBase struct {
	price    float64
	discount float64
	name     string
}

func (p *ProductBase) GetPrice() float64 {
	return p.price
}

func (p *ProductBase) SetPrice(price float64) {
	p.price = price
}

func (p *ProductBase) GetDiscount() float64 {
	return p.discount
}

func (p *ProductBase) SetDiscount(discount float64) {
	p.discount = discount
}

func (p *ProductBase) GetName() string {
	return p.name
}

func (p *ProductBase) SetName(name string) {
	p.name = name
}

// Функция для расчета стоимости списка товаров
func CalculateTotalPrice(products []Product) float64 {
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += product.GetPrice() * (1 - product.GetDiscount())
	}
	return totalPrice
}

func Lab7() {
	// Создание продуктов
	iphone := &ProductBase{price: 100, discount: 0.1, name: "Айфон"}
	ipad := &ProductBase{price: 200, discount: 0.2, name: "Айпад"}

	// Расчет стоимости списка товаров
	totalPrice := CalculateTotalPrice([]Product{iphone, ipad})
	fmt.Println("Общая цена:", totalPrice)
}

