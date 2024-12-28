package lab7

type Product interface {
	GetPrice() float64
	SetPrice(price float64)
	SetDiscount(discount float64)
}

func GetTotalPrice(ProductsList []Product) float64 {
	var defoultCost float64

	for _, product := range ProductsList {
		defoultCost += product.GetPrice()
	}
	return defoultCost
}
