package lab7

import (
	"fmt"
)

func Laba7() {
	var jeans = &Clothes{name: "джинсы", price: 3000, material: "полиэстер"}
	var chocolate = &Chocolate{name: "молочный шоколад", price: 70, brand: "Россия щедрая душа"}
	var comic = &Book{name: "Sailor Moon", price: 700, varieties: "манга"}

	comic.SetPrice(800)

	listProducts := []Product{jeans, chocolate, comic}
	fmt.Printf("Сумма товаров, без учета скидки, равна: %v рублей \n", CalculationSumProduct(listProducts))
	jeans.ApplyDiscount(10)
	chocolate.ApplyDiscount(5)
	comic.ApplyDiscount(15)
	fmt.Printf("Сумма товаров, c учетом скидки, равна: %v рублей \n", CalculationSumProduct(listProducts))
	fmt.Println(jeans)
	fmt.Println(chocolate)
	fmt.Println(comic)
}
