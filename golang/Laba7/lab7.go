package lab7

import (
	"fmt"
)

func CalculationSumProduct(listProducts []Product) string {
	var sum float64 = 0
	for _, product := range listProducts {
		sum += product.GetPrice()
	}
	s := fmt.Sprintf("%.2f", sum)
	return s
}

func Laba7() {
	var jeans Product = &Clothes{name: "джинсы", price: 3000, material: "полиэстер"}
	var chocolate Product = &Chocolate{name: "молочный шоколад", price: 70, brand: "Россия щедрая душа"}
	var comic Product = &Book{name: "Sailor Moon", price: 700, varieties: "манга"}

	if c, ok := jeans.(*Clothes); ok {
		c.SetMaterial("вискоза")
	} else {
		fmt.Println("Ошибка приведения типа для Clothes")
	}

	if ch, ok := chocolate.(*Chocolate); ok {
		ch.SetBrand("Alpen Gold")
	} else {
		fmt.Println("Ошибка приведения типа для Chocolate")
	}

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
