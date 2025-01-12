package lab7

import "fmt"

func GetTotalPrice(products []Product) float64 {
	var TotalPrice float64 = 0
	for _, product := range products {
		TotalPrice += product.GetPrice()
	}
	return TotalPrice
}

func Lab7() {
	Mobile := &Mobile{"Айфон 11 pro max", 50000, "Iphone", "gold"}
	Fruit := &Fruit{"Манго", 250, "свежее", "yellow"}
	Car := &Car{"Reno Logan", 500000, "Reno", "black"}
	products := []Product{Mobile, Fruit, Car}
	fmt.Println("Стоимость без скидок:", GetTotalPrice(products))

	Mobile.ApplyDiscount(25)
	Fruit.ApplyDiscount(40)
	Car.ApplyDiscount(10)
	fmt.Println("Стоимость после учёта скидки:", GetTotalPrice(products))
}