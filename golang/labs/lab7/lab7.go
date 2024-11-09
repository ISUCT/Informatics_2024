package labs

import "fmt"

func CalculateProductsSum(products []Product) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.getPrice()
	}
	return sum
}

func RunLab7() {
	product1 := &Food{50.15, "картошка", "пятерочка"}
	product2 := &Cosmetics{1700, "Chanel", "помада"}
	product3 := &Clothes{3000, "толстовка", "хлопок"}

	products := []Product{product1, product2, product3}
	fmt.Println("Общая стоимость товаров:", CalculateProductsSum(products), "рублей")
	product1.discount(20)
	product2.discount(10)
	product3.discount(15)
	fmt.Println("Общая стоимость товаров после применения скидок:", CalculateProductsSum(products), "рублей")
}
