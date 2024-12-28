package lab7

import "fmt"

func Lab7() {
	phone1 := Phone{
		Size:  100,
		Model: "Iphone",
		Price: 10000,
	}
	PC1 := PC{
		Model: "Model123",
		Price: 5000,
		RAM:   30,
	}
	Laptop1 := Laptop{
		Model: "ASUS",
		Price: 3000,
		ROM:   20,
	}
	ProductsList := []Product{&phone1, &PC1, &Laptop1}
	defoultCost := GetTotalPrice(ProductsList)
	fmt.Println("За товары вы заплатите:", defoultCost)
	for _, product := range ProductsList {
		product.SetDiscount(30)
	}

	defoultCost = GetTotalPrice(ProductsList)
	fmt.Println("С учетом всех скидок вы заплатите:", defoultCost)
	phone1.SetSize(660)
	PC1.SetRAM(50)
	Laptop1.SetROM(106)
}

func GetTotalPrice(ProductsList []Product) float64 {
	var defoultCost float64

	for _, product := range ProductsList {
		defoultCost += product.GetPrice()
	}
	return defoultCost
}
