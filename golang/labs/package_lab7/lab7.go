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
	product1 := &Food{50.15, "картошка", "вкусвилл", 17}
	product2 := &Cosmetics{"помада", 1700, "Chanel"}
	product3 := &Clothes{"толстовка", 3000, "Levi`s", "хлопок", "осень"}
	products := []Product{product1, product2, product3}
	fmt.Println("Продукты")
	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println("Общая стоимость товаров:", CalculateProductsSum(products), "рублей")
	product1.discount(20)
	product2.discount(10)
	product3.discount(15)
	fmt.Println("Общая стоимость товаров после применения скидок:", CalculateProductsSum(products), "рублей")
	season := product3.getSeason()
	material := product3.getMaterial()
	brand := product3.getBrand()
	fmt.Println("Информация о товаре")
	fmt.Println("материал:", material)
	fmt.Println("сезон:", season)
	fmt.Println("бренд:", brand)

}
