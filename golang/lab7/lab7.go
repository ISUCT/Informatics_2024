package lab7

import (
	"fmt"
)

type Product interface {
	GetName() string
	GetPrice() float64
	GetDescription() string
	ApplyDiscount(discount float64)
	SetPrice(newPrice float64)
	SetDescription(newDescription string)
	GetOldPrice() float64
}

func calculateTotalCost(products []Product) float64 {
	var totalCost float64
	for _, product := range products {
		totalCost += product.GetPrice()
	}
	return totalCost
}

func RunLab7() {
	book := &Book{
		Name:        "Война и мир",
		Price:       1300,
		Description: "роман",
		Author:      "Лев Николаевич Толстой",
	}

	phone := &ElectronicDevice{
		Name:        "iPhone 16",
		Price:       140000,
		Description: "новый смартфон",
		Brand:       "Apple",
		Model:       "iPhone 16",
	}

	products := []Product{book, phone}

	fmt.Printf("суммарная цена товвров: %.2f\n", calculateTotalCost(products))

	for _, product := range products {
		if _, ok := product.(*Book); ok {
			product.ApplyDiscount(0.10)
			fmt.Printf("Книга '%s': цена до скидки - %.2f, цена после скидки - %.2f\n", product.GetName(), product.GetOldPrice(), product.GetPrice())
		}
		if _, ok := product.(*ElectronicDevice); ok {
			product.ApplyDiscount(0.05)
			fmt.Printf("Телефон '%s': цена до скидки - %.2f, цена после скидки - %.2f\n", product.GetName(), product.GetOldPrice(), product.GetPrice())
		}
	}

	fmt.Printf("Цена после скидки: %.2f\n", calculateTotalCost(products))

	phone.SetPrice(125000)
	fmt.Printf("Телефон '%s': цена до скидки - %.2f, цена после скидки - %.2f\n", phone.GetName(), phone.GetOldPrice(), phone.GetPrice())

	book.SetDescription("самый популярный роман")
	fmt.Printf("изменённое описание книги: %s\n", book.GetDescription())
}
