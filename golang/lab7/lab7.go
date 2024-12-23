package lab7

import (
	"fmt"
)

func Lab7() {
	book1 := Book{name: "The Hitchhiker's Guide to the Galaxy", price: 15.99, numAuthors: 1}
	book2 := Book{name: "The Lord of the Rings", price: 30.99, numAuthors: 3}
	electronics1 := Electronics{name: "iPhone 13", price: 999.99, brand: "Apple", model: "iPhone 13", warranty: 12, isRefurbished: false}
	electronics2 := Electronics{name: "Samsung Galaxy S22", price: 799.99, brand: "Samsung", model: "Galaxy S22", warranty: 24, isRefurbished: true}

	totalBeforeDiscount := CalculateTotalPrice([]Product{&book1, &book2, &electronics1, &electronics2})
	fmt.Printf("Общая стоимость до скидок: %.2f\n", totalBeforeDiscount)

	book1.ApplyDiscount(0.1)
	electronics2.ApplyDiscount(0.2)

	totalAfterDiscount := CalculateTotalPrice([]Product{&book1, &book2, &electronics1, &electronics2})
	fmt.Printf("Общая стоимость после скидок: %.2f\n", totalAfterDiscount)
}
