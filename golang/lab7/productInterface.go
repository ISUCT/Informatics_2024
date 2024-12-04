package lab7

import "fmt"

type Product interface {
	SetDiscount(float64) error
	GetDiscount() float64
	GetPrice() float64
	GetName() string
}

func displayListProducts(PList []Product) {
	fmt.Println("____список товаров_____")
	for _, p := range PList {
		fmt.Printf(
			"________________________\nтовар \"%v\" стоит %v$ благодоря скидке в %v процентов\n",
			p.GetName(), p.GetPrice(), p.GetDiscount())
	}
}

func GetTotalPrice(ProductList []Product) float64 {
	var totalPrice float64

	for _, product := range ProductList {
		totalPrice += product.GetPrice()
	}
	return totalPrice
}
