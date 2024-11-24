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
	food := &Food{50.15, "картошка", "вкусвилл", 17}
	cosmetics := &Cosmetics{"помада", 1700, "Chanel"}
	clothes := &Clothes{"толстовка", 3000.99, "Levi`s", "хлопок", "осень"}
	products := []Product{food, cosmetics, clothes}
	fmt.Println("Товары")
	fmt.Println("Общая стоимость:", CalculateProductsSum(products), "рублей")
	for _, product := range products {
		fmt.Println(product.GetInfo())
	}

	for _, product := range products {
		product.setdiscount(20)
	}
	fmt.Println("Общая стоимость товаров после применения скидки 20%:", CalculateProductsSum(products), "рублей")

	fmt.Println("Информация про товар clothes")
	fmt.Println("название:", clothes.getName())
	fmt.Println("материал:", clothes.getMaterial())
	fmt.Println("сезон:", clothes.getSeason())
	fmt.Println("бренд:", clothes.getBrand())
}
