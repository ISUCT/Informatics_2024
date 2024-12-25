package Lab7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discount float64)
}

func Calculate(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func RunLab7() {
	shirt := &Shirt{
		Name: "Классическая рубашка", Price: 3000.00, Brand: "GUCCI", Color: "Белый", Material: "Хлопок", Size: "M",
	}
	bicycle := &Bicycle{
		Name: "Горный велосипед", Brand: "DOTA2", Price: 50000.00, BikeType: "Горный",
	}

	products := []Product{bicycle, shirt}
	fmt.Println("Общая стоимость ", Calculate(products))
	bicycle.ApplyDiscount(10)
	shirt.ApplyDiscount(20)

	fmt.Println("Общая стоимость товаров после применения скидок ", Calculate(products))

	fmt.Println("Тип велосипеда ", bicycle.BikeType)
	fmt.Println("Размер рубашки ", shirt.Size)

	bicycle.ChangeBikeType("Дорожный")
	shirt.ChangeColor("Синий")
	shirt.ChangeMaterial("Лен")
	shirt.ChangeSize("L")

	fmt.Println("Тип велосипеда изменен на ", bicycle.BikeType)
	fmt.Println("Цвет рубашки изменен на ", shirt.Color)
	fmt.Println("Материал рубашки изменен на ", shirt.Material)
	fmt.Println("Размер рубашки изменен на ", shirt.Size)
}
