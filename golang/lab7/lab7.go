package lab7

import "fmt"

type Product interface {
	GetInformation()
	GetPrice() float32
	Sale(float32)
	ChangePrice(float32)
	ChangeDescription(string)
}


func CalculateBue(list []Product) float32 {
	var sum float32 = 0
	for _, price := range list {
		sum += price.GetPrice()
	}
	return sum
}

func RunLab7Tasks() {
	var interface_cosmetic Product = &Cosmetic{"помада", "Gucci", 5000}
	var interface_techcic Product = &Technic{"ноутбук", "Apple", 169660}
	var interface_sweets Product = &Sweets{"Сникерс", "Солёная карамель", 85}
	interface_cosmetic.GetInformation()
	interface_techcic.GetInformation()
	interface_sweets.GetInformation()
	var bue []Product = []Product{interface_cosmetic, interface_techcic, interface_sweets}
	sum := CalculateBue(bue)
	fmt.Println("Товары стоят",sum)
	interface_techcic.ChangePrice(150000)
	interface_cosmetic.ChangeDescription("LV")
	interface_cosmetic.Sale(5)
	interface_sweets.Sale(60)
	interface_sweets.ChangeDescription("Белый шоколад")
	interface_techcic.GetInformation()
	interface_cosmetic.GetInformation()
	interface_sweets.GetInformation()

	sum = CalculateBue(bue)
	fmt.Println("Товары стоят",sum)
}