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
	potato := &Food{50.15, "картошка", "вкусвилл", 17}
	lipstick := &Cosmetics{"помада", 1700, "Chanel"}
	sweatshirt := &Clothes{"толстовка", 3000.99, "Levi`s", "хлопок", "осень"}
	products := []Product{potato, lipstick, sweatshirt}
	fmt.Println("Товары")
	fmt.Println("Общая стоимость:", CalculateProductsSum(products), "рублей")
	for _, product := range products {
		fmt.Println(product.GetInfo())
	}

	for _, product := range products {
		product.setdiscount(20)
	}
	fmt.Println("Общая стоимость товаров после применения скидки 20%:", CalculateProductsSum(products), "рублей")

	fmt.Println("Информация про товар толстовка")
	fmt.Println("материал:", sweatshirt.getMaterial())
	fmt.Println("сезон:", sweatshirt.getSeason())
	fmt.Println("бренд:", sweatshirt.getBrand())
}
