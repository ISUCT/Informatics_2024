package lab7

import (
	"fmt"
)


type Product interface{
	getName() string
	setName(string)
	getPrice() float64
	setPrice(float64)
	applyDicount(float64)
}


func ProductSum(productlist []Product) string{
	var sum float64 = 0
	for _, product := range productlist{
		sum += product.getPrice()
	}
	result := fmt.Sprintf("%r", sum)

}

func RunLab7(){
	pants := &Clothes{name: "штаны", price:1499, size: "M"}
	chocolate := &Food{name: "шоколад", price:99.99, calories:576}
	phone := &Food{name: "Iphone 15 Pro Max", price:170990, weight:227}

	pants.setSize("L")
	chocolate.setPrice(109.99)
	phone.setName("Iphone 16 Pro Max")
	phone.setWeight("225")

	productlist 
}