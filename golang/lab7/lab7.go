package lab7

import (
	"fmt"
)

type Product interface {
	ProductInfo() string
	ApplyDiscount(discount float64)
	GetPrice() float64
}

func CalculationSum(listproducts []Product) string {
	var sum float64 = 0
	for _, product := range listproducts {
		sum += product.GetPrice()
	}
	s := fmt.Sprintf("%.2f", sum)
	return s
}

func RunLab7() {
	var icetea Product = &Food{Name: "Холодный чай", Price: 70.37, Weight: 0.5}
	var bmw Product = &Car{Price: 20000000, Name: "BMW", Color: "Звёздная ночь", Model: "E60", Brand: "BMW"}
	var iphone Product = &Device{Price: 109000, Name: "IPhone", Color: "Титановый", Model: "15", Brand: "Apple"}

	fmt.Println(icetea.ProductInfo())
	fmt.Println(bmw.ProductInfo())
	fmt.Println(iphone.ProductInfo())

	listproducts := []Product{icetea, bmw, iphone}
	fmt.Printf("Сумма товаров, без учёта скидки, равна: %v рублей \n", CalculationSum(listproducts))
	icetea.ApplyDiscount(12)
	bmw.ApplyDiscount(18)
	iphone.ApplyDiscount(3)
	fmt.Printf("Сумма товаров, с учётом скидки, равна: %v рублей \n", CalculationSum(listproducts))
}
