package lab7

import "fmt"

func Lab7() {
	phone1 := Phone{
		Size:  80,
		Model: "CAMCUNG",
		Price: 10000,
	}
	PC1 := PC{
		Model: "VIPERPC",
		Price: 5000,
		RAM:   16,
	}
	Laptop1 := Laptop{
		Model: "MACDONALDOOK",
		Price: 3000,
		ROM:   16,
	}
	ProductsList := []Product{&phone1, &PC1, &Laptop1}
	defoultCost := GetTotalPrice(ProductsList)
	fmt.Println("Стоимость:", defoultCost)
	for _, product := range ProductsList {
		product.SetDiscount(30)
	}

	defoultCost = GetTotalPrice(ProductsList)
	fmt.Println("Стоимость с учётом скидок:", defoultCost)
	phone1.SetSize(700)
	PC1.SetRAM(60)
	Laptop1.SetROM(150)
}
