package lab7

import (
	"fmt"
)

type Products interface {
	GetName() string
	SetPrice(float64)
	GetPrice() float64
	ChangeData(string, float64, string)
	GetData()
	ApplyDiscount(float64)	
}

func calculateDiscount(products []Products) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.GetPrice()
	}

	return sum
}

func RunLab7() {
	product1 := NewBook("Harry Potter", "JK Rowling", 1000.0, "A book on Harry Potter")
	product2 := NewClothes("T-shirt", "Nike", 900.0, "Nice T-shirt")
	product3 := NewElectronics("Lamp", "Diode", 100.0, "Wonderful lamp")

	listOfProduct := []Products{product1, product2, product3}

	fmt.Printf("\nПродукты до изменения\n\n")
	product1.GetData()
	product2.GetData()
product1.ChangeData("Improved book Harry Potter", 1500.0, "Interesting book")
product2.ChangeData("Premium T-shirt", 1000.0, "A high-quality cotton T-shirt.")
	fmt.Printf("\nПродукты после изменения\n\n")
	product1.GetData()
	product2.GetData()

	fmt.Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%s-------%.2f $\n", product.GetName(), product.GetPrice())
	}
	fmt.Printf("Цена корзины до скидки: %.2f $\n\n", calculateDiscount(listOfProduct))

product1.ApplyDiscount(10)
product2.ApplyDiscount(20)
product3.ApplyDiscount(15)

	fmt. Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%s-------%.2f $\n", product.GetName(), product.GetPrice())
	}
	fmt.Printf("Цена корзины после скидки: %.2f $\n", calculateDiscount(listOfProduct))
}