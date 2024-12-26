package lab7

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

func Runlab7() {
	Product1 := &Phones{"iPhone 13", "128gb blue", 50000.0, "13 октября 2020 года"}
	Product2 := &Headphones{"Marshall Major IV", "Black", 7200.0, "14 октября 2020 года"}
	Product3 := &Laptop{"Huawei matepad 16", 1024, 130000.0, "AMD Ryzen 7 7840U"}
	ProductsWithoutDiscount := []Product{Product1, Product2, Product3}
	fmt.Printf("Цена без скидки: %2.f рублей\n", GetTotalPrice(ProductsWithoutDiscount))
	Product1.ApplyDiscount(5)
	Product2.ApplyDiscount(20)
	Product3.ApplyDiscount(25)
	fmt.Printf("Цена со скидкой: %2.f рублей\n", GetTotalPrice(ProductsWithoutDiscount))
	fmt.Println(Product1.getInfo())
	fmt.Println(Product2.getInfo())
	fmt.Println(Product3.getInfo())
	Product1.UpdateCharacteristics("256gb red")
	fmt.Println(Product1.getInfo())
	Product2.ChangeColor("Brown")
	fmt.Println(Product2.getInfo())
	Product3.ChangeCPU("i5 13500H")
	fmt.Println(Product3.getInfo())
}
