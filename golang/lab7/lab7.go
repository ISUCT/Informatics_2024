package lab7

import "fmt"

type Products interface {
	GetInformation()
	GetPrice() float32
	ApplyDiscount(float32)
	ChangePrice(float32)
	ChangeDescription(string)
}


func CalculateBuy(list []Products) float32 {
	var sum float32 = 0
	for _, price := range list {
		sum += price.GetPrice()
	}
	return sum
}

func RunLab7Tasks() {
	var cosmetic Products = &Cosmetic{"помада", "Gucci", 5000}
	var techcic Products = &Technic{"ноутбук", "Apple", 169660}
	var sweets Products = &Sweets{"Сникерс", "Солёная карамель", 85}
	cosmetic.GetInformation()
	techcic.GetInformation()
	sweets.GetInformation()
	var buy []Products = []Products{cosmetic , techcic, sweets}
	sum := CalculateBuy(buy)
	fmt.Println("Товары стоят",sum)
	techcic.ChangePrice(150000)
	cosmetic.ChangeDescription("LV")
	cosmetic.ApplyDiscount(5)
	sweets.ApplyDiscount(60)
	sweets.ChangeDescription("Белый шоколад")
	techcic.GetInformation()
	cosmetic.GetInformation()
	sweets.GetInformation()

	sum = CalculateBuy(buy)
	fmt.Println("Товары стоят",sum)
}
