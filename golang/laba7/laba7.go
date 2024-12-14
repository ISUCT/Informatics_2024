package laba7

import "fmt"

type Product interface {
	ApplyDiscount(discount float64)
	SetPrice(newPrice float64)
	GetPrice() float64
	getInfo() string
}

func GetTotalPrice(products []Product) float64 {
	var TotalPrice float64 = 0
	for _, Product := range products {
		TotalPrice += Product.GetPrice()
	}
	return TotalPrice
}

func RunLaba7() {
	Product1 := &Gaming{"Keyboard", "Cherry mx red", 5000.0, 2015.0}
	Product2 := &Mobiles{"Nokia", "White-Blue", 500.0, 2001.0}
	Product3 := &VideoGames{"MafiaII", 15, 500.0, "2K"}
	ProductsWithoutDiscount := []Product{Product1, Product2, Product3}
	fmt.Printf("Цена без скидки: %2.f рублей\n", GetTotalPrice(ProductsWithoutDiscount))
	Product1.ApplyDiscount(10)
	Product2.ApplyDiscount(20)
	Product3.ApplyDiscount(25)
	fmt.Printf("Цена со скидкой: %2.f рублей\n", GetTotalPrice(ProductsWithoutDiscount))
	fmt.Println(Product1.getInfo())
	fmt.Println(Product2.getInfo())
	fmt.Println(Product3.getInfo())
	Product1.UpdateCharacteristics("Cherry mx silver speed")
	fmt.Println(Product1.getInfo())
	Product2.ChangeColor("Red")
	fmt.Println(Product2.getInfo())
	Product3.UpdateCompany("Valve")
	fmt.Println(Product3.getInfo())
}
